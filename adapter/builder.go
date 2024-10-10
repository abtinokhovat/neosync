package adapter

import (
	"neosync/adapter/mariadb"
	"neosync/internal/config"
)

type Adapters struct {
	MariaDB *mariadb.DB
}

func Build(cfg *config.Config) *Adapters {
	return &Adapters{
		MariaDB: mariadb.New(cfg.DB),
	}
}
