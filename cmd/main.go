package main

import (
	"neosync/delivery/http"
	"neosync/internal/config"
	"neosync/internal/logger"
	"neosync/pkg/buildinfo"
)

func main() {
	// print build info
	buildinfo.Print()
	// config service
	cfg := config.C()
	// logger service
	logger.L()

	server := http.NewServer(cfg.Server, nil)
	server.Start()
}
