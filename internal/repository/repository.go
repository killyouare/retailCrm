package repository

import (
	"context"
	"crmtest/internal/domain/order"
	"crmtest/internal/pkg/db"
	"crmtest/internal/pkg/util"
)

type Repository interface {
	Migrate(ctx context.Context) error
	CreateOrders(ctx context.Context, orders []order.Order) error
}

type repository struct {
	logger     logger.Logger
	connection db.Connection
	table      string
}

func New(logger logger.Logger, connection db.Connection, table string) (Repository, error) {
	return &repository{
		logger:     logger,
		connection: connection,
		table:      table,
	}, nil
}
