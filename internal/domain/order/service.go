package order

import "context"

type Repository interface {
	UpdateStatus(ctx context.Context, id uint, status Status) error // UpdateManyStatuses(ctx context.Context, req []map[uint]Status) error
	GetPendingOrders(ctx context.Context) ([]Order, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}
