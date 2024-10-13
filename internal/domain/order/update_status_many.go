package order

import (
	"context"
	"neosync/internal/domain/provider"
)

type UpdateStatusManyRequest struct {
	TrackingStatusMapping map[string]provider.AdapterResponseItem
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
		if Status(providerOrderStatus.Status) == order.Status {
			continue
		}

		// the order status was changed and should be updated here
		rErr := s.repo.UpdateStatus(ctx, order.ID, Status(providerOrderStatus.Status))
		if rErr != nil {
			// retry policy here
		}
	}

	return nil
}
