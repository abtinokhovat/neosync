package config

import "neosync/delivery/http"

func defaultConfig() *Config {
	return &Config{
		Server: http.Config{
			EnableBanner: false,
			Port:         8080,
			Timeout:      600000,
			MetricPort:   8084,
		},
	}
}
