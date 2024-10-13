package main

import (
	"neosync/delivery/http"
	"neosync/delivery/http/providerhandler"
	"neosync/internal/config"
	"neosync/internal/domain/provider"
	"neosync/internal/infra/adapter"
	"neosync/internal/infra/db/mariadb"
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

	adapters := adapter.Build(cfg)
	maria := mariadb.Builder(adapters.MariaDB, adapters.Gorm)
	providerService := provider.NewService(maria.Provider, adapters.OperationProviders)

	routers := []http.Router{
		providerhandler.New(providerService),
	}

	server := http.NewServer(cfg.Server, routers)
	server.Start()
}
