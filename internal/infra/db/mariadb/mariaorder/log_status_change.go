package mariaorder

import (
	"context"
	"log/slog"
	"neosync/internal/domain/order"
	"neosync/internal/logger"
	"neosync/pkg/richerror"
	"time"
)

func (d *DB) LogStatusChange(ctx context.Context, orderID uint, status string) error {
	const op = "mariaorder.LogStatusChange"

	// Define the SQL statement for inserting a status change record
	query := "insert into `order_status_history` (order_id, status, changed_at) value (?, ?, ?)"

	// Execute the query with parameters
	_, err := d.conn.Conn().ExecContext(ctx, query, orderID, status, time.Now())
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}

// UpdateStatusAndLogChange both update the order status and log the change with a transaction
func (d *DB) UpdateStatusAndLogChange(ctx context.Context, orderID uint, status order.Status) error {
	const op = "mariaorder.updateStatusAndLogChange"

	// begin a new transaction
	tx, err := d.conn.Conn().BeginTx(ctx, nil)
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	// rollback the transaction in case of failure
	defer func() {
		if p := recover(); p != nil {
			rErr := tx.Rollback()
			if rErr != nil {
				logger.L().Error(op, slog.Any("transaction", tx), slog.Any("rollback error", rErr))
				return
			}
		} else if err != nil {
			rErr := tx.Rollback()
			if rErr != nil {
				logger.L().Error(op, slog.Any("transaction after rollback", tx), slog.Any("rollback error", rErr))
				return
			}
		} else {
			err = tx.Commit()
		}
	}()

	// step 1: update the order status
	const updateQuery = "update orders set status = ? where id = ?"
	_, err = tx.ExecContext(ctx, updateQuery, status, orderID)
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	// step 2: log the status change
	const logQuery = "insert into order_status_history (order_id, status, changed_at) values (?, ?, ?)"
	_, err = tx.ExecContext(ctx, logQuery, orderID, status, time.Now())
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}
