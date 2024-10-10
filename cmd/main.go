package main

import (
	"neosync/delivery/http"
	"neosync/internal/config"
	"neosync/internal/logger"
)

func main() {
	// config service
	cfg := config.C()
	// logger service
	logger.L()

	server := http.NewServer(cfg.Server, nil)
	server.Start()
}
