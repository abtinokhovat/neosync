package config

import (
	"neosync/adapter/mariadbadapter"
	"neosync/delivery/http"
)

type Config struct {
	Server http.Config           `koanf:"server"`
	DB     mariadbadapter.Config `koanf:"db"`
}
