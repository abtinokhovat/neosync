package customer

import "errors"

var (
	ErrMsCustomerNotFound = errors.New("customer not found")
)

type Customer struct {
	ID          uint
	Name        string
	PhoneNumber string
}
