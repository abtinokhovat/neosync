package config

import (
	"neosync/adapter/mariadbadapter"
	"neosync/delivery/http"
	"neosync/pkg/migrator"
)

type Config struct {
	Server   http.Config           `koanf:"server"`
	DB       mariadbadapter.Config `koanf:"db"`
	Migrator migrator.Config       `koanf:"migrator"`
}
