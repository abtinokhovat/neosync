package mariaorder

import (
	"neosync/internal/domain/order"
	"neosync/internal/infra/adapter/mariadb"
)

type DB struct {
	conn *mariadb.DB
}

func New(conn *mariadb.DB) *DB {
	return &DB{
		conn: conn,
	}
}

func scanOrder(scanner mariadb.Scanner) (order.Order, error) {
	o := order.Order{}
	err := scanner.Scan(
		&o.ID, &o.CreatedAt,
		&o.UpdatedAt, &o.Status,
		&o.TrackingCode, &o.CustomerID,
		&o.ProviderID,
	)
	return o, err
}
