package config

import (
	"neosync/internal/infra/cron"
)

func defaultConfig() *Config {
	return &Config{
		OrderUpdater: cron.OrderUpdaterConfig{
			CronExpression: "0 0 * * *",
			TimeoutMinutes: 10,
		},
	}
}
