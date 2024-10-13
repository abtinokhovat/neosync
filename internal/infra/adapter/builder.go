package adapter

import (
	"neosync/internal/config"
	"neosync/internal/domain/provider"
	"neosync/internal/infra/adapter/gorm"
	"neosync/internal/infra/adapter/mariadb"
	"neosync/internal/infra/adapter/providermock1"
)

type Adapters struct {
	MariaDB            *mariadb.DB
	Gorm               *gorm.DB
	OperationProviders map[string]provider.Adapter
}

func Build(cfg *config.Config) *Adapters {
	mockAdapter1 := providermock1.New()

	return &Adapters{
		MariaDB: mariadb.New(cfg.DB),
		Gorm:    gorm.New(cfg.DB),
		OperationProviders: map[string]provider.Adapter{
			mockAdapter1.Name(): mockAdapter1,
		},
	}
}
