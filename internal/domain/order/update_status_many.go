package order

import (
	"context"
	"fmt"
	"neosync/internal/domain/provider"
)

var orderStatusPickedUpMessage = "your order (%s) was picked up by the provider"

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
		rErr := s.repo.UpdateStatusAndLogChange(ctx, order.ID, Status(providerOrderStatus.Status))
		if rErr != nil {
			// retry policy here
		}

		go func() {
			// on the ProviderSeen -> PickedUp status change sending a notification
			if !(Status(providerOrderStatus.Status) == PickedUp && order.Status == ProviderSeen) {
				return
			}

			snErr := s.notifier.SendNotification(ctx, fmt.Sprintf(orderStatusPickedUpMessage, order.TrackingCode), order.CustomerID)
			if snErr != nil {
				// retry policy here
				// or adding a table to store the sending status and create another batch job over that to periodically retry them
			}
		}()

	}

	return nil
}
