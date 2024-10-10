package config

import (
	"neosync/adapter/mariadb"
	"neosync/delivery/http"
	"neosync/pkg/migrator"
)

type Config struct {
	Server   http.Config     `koanf:"server"`
	DB       mariadb.Config  `koanf:"db"`
	Migrator migrator.Config `koanf:"migrator"`
}
