package order

import (
	"context"
)

type UpdateStatusManyRequest struct {
	TrackingStatusMapping map[string]Status
}

func (s Service) UpdateStatusMany(ctx context.Context, req UpdateStatusManyRequest) error {
	orders, err := s.GetPendingForReview(ctx)
	if err != nil {
		return err
	}

	// updating each orders
	for _, order := range orders {
		providerOrderStatus, has := req.TrackingStatusMapping[order.TrackingCode]
		// if the providers did not have the order data
		if !has {
			continue
		}

		// if the providers had the order but the status did not change from the last check
		if providerOrderStatus == order.Status {
			continue
		}

		// the order status was changed and should be updated here
		rErr := s.repo.UpdateStatus(ctx, order.ID, providerOrderStatus)
		if rErr != nil {
			// retry policy here
		}
	}

	return nil
}
