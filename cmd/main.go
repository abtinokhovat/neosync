package main

import (
	"neosync/delivery/http"
	"neosync/internal/config"
	"neosync/internal/infra/db/mariadb/migrate"
	"neosync/internal/logger"
	"neosync/pkg/buildinfo"
	"neosync/pkg/migrator"
)

func main() {
	// print build info
	buildinfo.Print()
	// config service
	cfg := config.C()
	// logger service
	logger.L()

	// migrate database
	mgr := migrator.New(cfg.Migrator, cfg.DB.String(), migrate.Provide())
	mgr.Up()

	//adapters := adapter.Build(cfg)
	//dbs := mariadb.Builder(adapters.MariaDB)
	//
	////_ = order.NewService(databases.Order)
	//pv := provider.NewService(dbs.Provider, adapters.OperationProviders)

	server := http.NewServer(cfg.Server, nil)
	server.Start()
}
