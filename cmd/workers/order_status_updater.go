package main

import (
	"neosync/internal/config"
	"neosync/internal/domain/order"
	"neosync/internal/domain/provider"
	"neosync/internal/infra/adapter"
	"neosync/internal/infra/cron"
	"neosync/internal/infra/db/mariadb"
	"neosync/internal/logger"
)

func main() {
	cfg := config.C()
	logger.L()

	adapters := adapter.Build(cfg)
	databases := mariadb.Builder(adapters.MariaDB)

	orderService := order.NewService(databases.Order)
	providerService := provider.NewService(databases.Provider, adapters.OperationProviders)

	// create the updater and start the cron job
	updater := cron.NewOrderStatusUpdater(cfg.OrderUpdater, orderService, providerService)
	updater.StartCronJob()

	// putting a select statement here to not let the binary execute after each cron execution
	select {}
}
