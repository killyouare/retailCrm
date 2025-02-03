package repository

import (
	"context"
	"crmtest/internal/domain/order"
	"database/sql"
	"fmt"
	"strings"
)

var (
	dateLayout = "2006-01-02 15:04:05"
)

var (
	failExecute    = "failed to execute query: %w"
	failBuildQuery = "failed to build query: %w"
	failPrepare    = "failed to prepare query: %w"
	failCommitTx   = "failed to commit transaction: %w"
	failBeginTx    = "failed to begin transaction: %w"
)

func (r *repository) CreateOrders(ctx context.Context, orders []order.Order) error {
	tx, err := r.createTransaction(ctx)
	if err != nil {
		return fmt.Errorf(failBeginTx, err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	query, args, err := r.buildInsertQuery(orders)
	if err != nil {
		return fmt.Errorf(failBuildQuery, err)
	}
	stmt, err := tx.PrepareContext(ctx, query)

	if err != nil {
		return fmt.Errorf(failPrepare, err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return fmt.Errorf(failExecute, err)
	}

	err = tx.Commit()

	if err != nil {
		return fmt.Errorf(failCommitTx, err)
	}

	r.logger.Info("Заказы загружены в базу")

	return nil
}

func (r *repository) createTransaction(ctx context.Context) (*sql.Tx, error) {
	tx, err := r.connection.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (r *repository) buildInsertQuery(orders []order.Order) (string, []interface{}, error) {
	baseQuery := fmt.Sprintf("INSERT INTO %s (order_id, client_id, order_number, fakt_data, totalSumm, prepaySum) VALUES ", r.table)

	var values []string
	var args []interface{}

	for _, o := range orders {
		values = append(values, "(?, ?, ?, ?, ?, ?)")
		o, err := orderToSql(&o)
		if err != nil {
			return "", nil, err
		}
		args = append(args, o.orderId, o.clientId, o.number, o.created, o.totalSum, o.prepaySum)
	}

	duplicateKey := ` ON DUPLICATE KEY UPDATE  
	client_id = VALUES(client_id),
    order_number = VALUES(order_number),
    fakt_data = VALUES(fakt_data),
    podtverzden_otpravka = VALUES(podtverzden_otpravka),
    totalSumm = VALUES(totalSumm),
    prepaySum = VALUES(prepaySum)`
	query := baseQuery + strings.Join(values, ", ") + duplicateKey
	return query, args, nil
}

type sqlOrder struct {
	clientId  int
	orderId   int
	number    string
	created   string
	totalSum  int64
	prepaySum int64
}

func orderToSql(o *order.Order) (*sqlOrder, error) {
	return &sqlOrder{
		clientId:  o.ClientId,
		orderId:   o.OrderId,
		number:    o.Number,
		created:   o.CreatedAt.Format(dateLayout),
		totalSum:  o.TotalSum,
		prepaySum: o.PrepaySum,
	}, nil
}

//func sqlOrderToOrder(s *sqlOrder) *order.Order {
//	//TODO: реализовать
//	panic("implement me")
//}
