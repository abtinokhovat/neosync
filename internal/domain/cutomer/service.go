package cutomer

import "context"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s Service) GetPhoneNumber(ctx context.Context, userID uint) (string, error) {
	// this is a mock implementation here and should be read the customer table from the database and return the correct customer number
	return "09121111111", nil
}
