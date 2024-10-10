package main

import (
	"neosync/internal/config"
	"neosync/internal/logger"
)

func main() {
	// config service
	config.C()
	// logger service
	logger.L()
}
