package db

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"time"
)

type Connection interface {
	Close() error
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	sqlx.ExecerContext
}

type connection struct {
	db *sqlx.DB
}

func New(dsn string) (Connection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := sqlx.ConnectContext(ctx, "mysql", dsn)

	if err != nil {
		return nil, err
	}

	conn.SetConnMaxLifetime(time.Minute * 30)
	conn.SetMaxOpenConns(10)

	return &connection{
		db: conn,
	}, nil
}

func (c *connection) Close() error {
	return c.db.Close()
}

func (c *connection) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return c.db.BeginTx(ctx, opts)
}

func (c *connection) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return c.db.ExecContext(ctx, query, args...)
}
