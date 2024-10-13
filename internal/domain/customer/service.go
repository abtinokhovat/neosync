package customer

import (
	"context"
)

type Repository interface {
	Get(ctx context.Context, customerID uint) (Customer, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}
