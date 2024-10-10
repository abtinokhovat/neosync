package adapter

import (
	"neosync/adapter/mariadbadapter"
	"neosync/internal/config"
)

type Adapters struct {
	MariaDB *mariadbadapter.DB
}

func Build(cfg *config.Config) *Adapters {
	return &Adapters{
		MariaDB: mariadbadapter.New(cfg.DB),
	}
}
