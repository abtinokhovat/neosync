package provider

import (
	"context"
	"errors"
	"neosync/internal/domain/order"
)

var (
	ErrMsgWrongProviderNotFound = errors.New("provider not found")
)

type Adapter interface {
	// Name is a getter for a provider, so we could identify it and bing it on the correct requests
	// I could have identified by their ids in db but this may change, so I chose to identify them by their name
	Name() string
	// GetAll fetch all the orders from the providers by their status and tracking number
	GetAll(ctx context.Context) (map[string]order.Status, error)
}

type Repository interface {
	GetAll(ctx context.Context) (map[uint]Provider, error)
}

type Service struct {
	repo     Repository
	adapters map[string]Adapter
}

func NewService(repo Repository, adapters map[string]Adapter) *Service {
	return &Service{
		repo:     repo,
		adapters: adapters,
	}
}
