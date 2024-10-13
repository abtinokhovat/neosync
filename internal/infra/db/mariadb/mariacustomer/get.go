package mariacustomer

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"neosync/internal/domain/customer"
	"neosync/pkg/richerror"
)

func (db *DB) Get(ctx context.Context, customerID uint) (customer.Customer, error) {
	const op = "mariacustomer.Get"
	var c customer.Customer

	if err := db.conn.Conn().WithContext(ctx).First(&c, customerID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer.Customer{}, richerror.New(op).WithErr(customer.ErrMsCustomerNotFound).WithKind(richerror.KindNotFound)
		}
		return customer.Customer{}, err
	}

	return c, nil
}
