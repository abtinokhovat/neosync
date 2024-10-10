package main

import (
	"neosync/delivery/http"
	"neosync/internal/config"
	"neosync/internal/dependency/adapter"
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

	_ = adapter.Build(cfg)

	server := http.NewServer(cfg.Server, nil)
	server.Start()
}
