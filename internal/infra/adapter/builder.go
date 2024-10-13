package adapter

import (
	"neosync/internal/config"
	"neosync/internal/domain/provider"
	"neosync/internal/infra/adapter/mariadb"
)

type Adapters struct {
	MariaDB            *mariadb.DB
	OperationProviders map[string]provider.Adapter
}

func Build(cfg *config.Config) *Adapters {
	return &Adapters{
		MariaDB:            mariadb.New(cfg.DB),
		OperationProviders: map[string]provider.Adapter{},
	}
}
