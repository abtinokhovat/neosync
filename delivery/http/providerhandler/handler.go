package providerhandler

import "neosync/internal/domain/provider"

type Handler struct {
	service *provider.Service
}

func New(service *provider.Service) *Handler {
	return &Handler{
		service: service,
	}
}
