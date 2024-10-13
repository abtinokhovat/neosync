package provider

import (
	"context"
	"log/slog"
	"maps"
	"neosync/internal/logger"
	"neosync/pkg/richerror"
	"sync"
)

type GetAllRequest struct {
	ProviderIDs []uint `json:"provider_ids"`
}

type GetAllResponse struct {
	Mapping map[string]AdapterResponseItem
}

// BatchFetchAll is function to fetch data from all requested providers
func (s Service) BatchFetchAll(ctx context.Context, req GetAllRequest) (GetAllResponse, error) {
	const op = "provider.BatchFetchAll"

	if req.ProviderIDs == nil || len(req.ProviderIDs) == 0 {
		return GetAllResponse{}, richerror.New("no provider id provided")
	}

	// fetching order statuses from all providers
	trackingCodeStatusMap := make(map[string]AdapterResponseItem)

	// getting all providers to search on it by their names
	providers, err := s.GetAll(ctx)
	if err != nil {
		return GetAllResponse{}, err
	}

	adapters := filterAdapters(req.ProviderIDs, s.adapters, providers)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for _, adapter := range adapters {
		wg.Add(1)
		go func(adapter Adapter) {
			defer wg.Done()
			codes, pErr := adapter.GetAll(ctx)
			if pErr != nil {
				// retry policy here
			}

			// merge maps
			mu.Lock()
			maps.Copy(trackingCodeStatusMap, codes)
			mu.Unlock()
		}(adapter)
	}
	wg.Wait()

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
