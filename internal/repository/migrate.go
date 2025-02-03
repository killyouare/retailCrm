package repository

import (
	"context"
	"fmt"
)

func (r *repository) Migrate(ctx context.Context) error {
	createTable := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ", r.table)
	//TODO total sum > 0? добавить check
	//order_id использован как ключ, для ON DUPLICATE KEY UPDATE
	//не совсем правильный подход, должен быть системный айдишник первичным
	//плюс лучше логгировать каждое изменение заказа
	// странно получилось по названию полей, каждое в раных кейсах, в идеале такого не допускать
	tableStruct := `(
		order_id int PRIMARY KEY,
    	client_id int NOT NULL,
		order_number VARCHAR(15) NOT NULL,
    	fakt_data DATETIME NOT NULL,
    	podtverzden_otpravka DATETIME NULL,
    	totalSumm BIGINT NOT NULL,
    	prepaySum BIGINT NOT NULL
)`
	query := createTable + tableStruct + ";"
	_, err := r.connection.ExecContext(ctx, query)
	return err

}
