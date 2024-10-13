package notifer

import (
	"context"
)

type RecipientProvider interface {
	GetPhoneNumber(ctx context.Context, userID uint) (string, error)
}

type Service struct {
	recipient RecipientProvider
}

func NewService(recipient RecipientProvider) *Service {
	return &Service{
		recipient: recipient,
	}
}
