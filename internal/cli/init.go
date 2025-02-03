package cli

import (
	"context"
	"crmtest/config"
	"crmtest/internal/action"
	"crmtest/internal/pkg/db"
	http2 "crmtest/internal/pkg/http"
	"crmtest/internal/pkg/retail"
	logger "crmtest/internal/pkg/util"
	"crmtest/internal/repository"
	"crmtest/internal/service"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
	"os/signal"
)

var (
	failMigrate = "failed to migrate db: %w"
)

func (a *App) initLogger() error {
	log, err := logger.New()

	if err != nil {
		return err
	}

	a.logger = log

	return nil
}

func (a *App) initConfig() error {
	//TODO: путь захардкожен
	c, err := config.New("./.env")
	if err != nil {
		return err
	}

	a.config = c
	return nil
}

func (a *App) initHttp() error {
	httpClient, err := http2.New(&http.Client{Timeout: a.config.Http.Timeout})

	if err != nil {
		return err
	}

	a.http = httpClient
	return nil
}

func (a *App) initRetail() error {
	client, err := retail.New(a.http, a.logger, a.config.Api.Url, a.config.Api.Apikey)
	if err != nil {
		return err
	}
	a.retail = client
	return nil
}

func (a *App) initDB() error {
	dbConn, err := db.New(a.config.Database.Dsn)

	if err != nil {
		return err
	}

	a.closers = append(a.closers, func() error {
		return dbConn.Close()
	})

	a.dbConn = dbConn

	return nil
}

func (a *App) initRepository() error {
	r, err := repository.New(a.logger, a.dbConn, a.config.Database.Table)

	if err != nil {
		return err
	}

	a.repository = r

	if err = a.repository.Migrate(context.Background()); err != nil {
		return fmt.Errorf(failMigrate, err)
	}
	return nil
}

func (a *App) initShutdown() error {
	a.closeCh = make(chan os.Signal, 1)
	signal.Notify(a.closeCh, os.Interrupt)

	return nil
}

func (a *App) initService() error {
	getOrderService, err := service.NewGetOrderService(a.logger, a.retail)

	if err != nil {
		return err
	}

	a.service = getOrderService
	return nil
}

func (a *App) initAction() error {
	getOrderAction, err := action.NewGetOrdersAction(a.logger, a.repository, a.service)
	if err != nil {
		return err
	}
	a.getOrderAction = getOrderAction
	return nil
}
