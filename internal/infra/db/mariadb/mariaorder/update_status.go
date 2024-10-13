package mariaorder

import (
	"context"
	"neosync/internal/domain/order"
	"neosync/pkg/richerror"
)

func (d *DB) UpdateStatus(ctx context.Context, id uint, status order.Status) error {
	const op = "mariaorder.UpdateStatus"
	const updateQuery = "update orders set status = ? where id = ?"

	_, err := d.conn.Conn().ExecContext(ctx, updateQuery, status, id)
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}
