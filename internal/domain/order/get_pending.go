package order

import "context"

func (s Service) GetPendingForReview(ctx context.Context) ([]Order, error) {
	orders, err := s.repo.GetPendingOrders(ctx)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
