package provider

import (
	"context"
	"log/slog"
	"maps"
	"neosync/internal/domain/order"
	"neosync/internal/logger"
)

type GetAllRequest struct {
	ProviderIDs []uint `json:"provider_ids"`
}

type GetAllResponse struct {
	Mapping map[string]order.Status
}

// BatchFetchAll is function to fetch data from all requested providers
func (s Service) BatchFetchAll(ctx context.Context, req GetAllRequest) (GetAllResponse, error) {
	// fetching order statuses from all providers
	trackingCodeStatusMap := make(map[string]order.Status)

	// getting all providers to search on it by their names
	providers, err := s.repo.GetAll(ctx)
	if err != nil {
		return GetAllResponse{}, err
	}

	adapters := filterAdapters(req.ProviderIDs, s.adapters, providers)

	// TODO: make this concurrent
	for _, adapter := range adapters {
		codes, pErr := adapter.GetAll(ctx)
		if pErr != nil {
			// retry policy here
		}

		// merge maps
		maps.Copy(trackingCodeStatusMap, codes)
	}

	return GetAllResponse{Mapping: trackingCodeStatusMap}, nil
}

// filterAdapters filters the adapter by the provider providerIDs
func filterAdapters(providerIDs []uint, adapters map[string]Adapter, providers map[uint]Provider) []Adapter {
	const op = "provider.filterAdapters"

	filteredAdapters := make([]Adapter, 0)

	for _, id := range providerIDs {
		provider, ok := providers[id]
		if !ok {
			// wrong id was passed here
			// we can return error here, but I just added log to do not interrupt the flow
			// obviously if anything went wrong here is the programmer bug so it can be fixed with the correct logic
			logger.L().Warn(op, slog.Any("message", ErrMsgWrongProviderNotFound))
			continue
		}

		// appending the found adapter and making it ready for the batch request
		filteredAdapters = append(filteredAdapters, adapters[provider.Name])
	}

	return filteredAdapters
}
