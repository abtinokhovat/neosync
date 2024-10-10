package config

import "neosync/delivery/http"

type Config struct {
	Server http.Config `koanf:"server"`
}
