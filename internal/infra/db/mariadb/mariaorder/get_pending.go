package mariaorder

import (
	"context"
	"database/sql"
	"errors"
	"neosync/internal/domain/order"
	"neosync/pkg/richerror"
)

func (d *DB) GetPendingOrders(ctx context.Context) ([]order.Order, error) {
	const op = "mariadb.GetPendingOrders"
	const query = "select * from orders where status > 1 and status < 5"

	rows, err := d.conn.Conn().QueryContext(ctx, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []order.Order{}, nil
		}

		return nil, richerror.New(op).WithKind(richerror.KindUnexpected).WithErr(err)
	}

	orders := make([]order.Order, 0)
	for rows.Next() {
		scannedOrder, sErr := scanOrder(rows)
		if sErr != nil {
			return nil, richerror.New(op).WithKind(richerror.KindUnexpected).WithErr(sErr)
		}

		orders = append(orders, scannedOrder)
	}

	return orders, nil
}
