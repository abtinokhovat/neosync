package order

import "context"

type Repository interface {
	UpdateStatusAndLogChange(ctx context.Context, orderID uint, status Status) error
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
