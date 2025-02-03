package cli

import (
	"context"
	"crmtest/config"
	"crmtest/internal/action"
	"crmtest/internal/command"
	"crmtest/internal/pkg/db"
	http2 "crmtest/internal/pkg/http"
	"crmtest/internal/pkg/retail"
	"crmtest/internal/pkg/util"
	"crmtest/internal/repository"
	"crmtest/internal/service"
	"os"
)

type App struct {
	retail         retail.Client
	config         *config.Config
	http           http2.Client
	logger         logger.Logger
	dbConn         db.Connection
	service        service.Service
	getOrderAction action.Action

	repository repository.Repository
	closers    []func() error
	closeCh    chan os.Signal
}

func New() (*App, error) {
	var err error
	app := &App{}

	if err = app.initConfig(); err != nil {
		return nil, err
	}

	if err = app.initLogger(); err != nil {
		return nil, err
	}

	if err = app.initHttp(); err != nil {
		return nil, err
	}

	if err = app.initRetail(); err != nil {
		return nil, err
	}

	if err = app.initDB(); err != nil {
		return nil, err
	}

	if err = app.initRepository(); err != nil {
		return nil, err
	}

	if err = app.initService(); err != nil {
		return nil, err
	}

	if err = app.initAction(); err != nil {
		return nil, err
	}

	if err = app.initShutdown(); err != nil {
		return nil, err
	}

	return app, nil

}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd, err := command.Parse(&command.OrderDep{
		A: a.getOrderAction,
	}, os.Args...)
	if err != nil {
		return err
	}

	//Закрыть приложение на сигнале Interrupt
	go func() {
		<-a.closeCh
		cancel()
	}()
	//Закрытие всех ресурсов, стек, поэтому с конца перебор
	defer func() {
		for i := len(a.closers) - 1; i >= 0; i-- {
			err := a.closers[i]()
			if err != nil {
				a.logger.Error("failed to close resource", "i", i, "error", err)
			}
		}
	}()

	err = cmd.Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
