package order

import "context"

type Repository interface {
	UpdateStatus(ctx context.Context, id uint, status Status) error // UpdateManyStatuses(ctx context.Context, req []map[uint]Status) error
	GetPendingOrders(ctx context.Context) ([]Order, error)
}

type Notifier interface {
	SendNotification(ctx context.Context, message string, userID uint) error
}

type Service struct {
	repo     Repository
	notifier Notifier
}

func NewService(repo Repository, notifier Notifier) *Service {
	return &Service{repo, notifier}
}
