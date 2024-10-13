package config

import (
	"neosync/delivery/http"
	"neosync/internal/infra/adapter/mariadb"
	"neosync/internal/infra/cron"
	"neosync/pkg/migrator"
)

type Config struct {
	Server       http.Config             `koanf:"server"`
	DB           mariadb.Config          `koanf:"db"`
	Migrator     migrator.Config         `koanf:"migrator"`
	OrderUpdater cron.OrderUpdaterConfig `koanf:"order_updater"`
}
