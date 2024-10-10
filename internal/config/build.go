package config

import (
	"log"
	"neosync/pkg/config"
	"os"
)

const (
	prefix       = "NEOSYNC"
	delimiter    = "."
	envDelimiter = "__"
)

// making a
var loader *config.Loader[Config]

// init initialize the config load module to load the config from various sources in the application
func init() {
	envName := "NEOSYNC_CONFIG_PATH"
	fallbackPath := "./deploy/config.yml"

	// get YAML file path from environment variable or use fallback
	yamlPath := os.Getenv(envName)
	if yamlPath == "" {
		yamlPath = fallbackPath
	}

	opt := config.Option{
		Prefix:       prefix,
		Delimiter:    delimiter,
		EnvDelimiter: envDelimiter,
		YamlFilePath: yamlPath,
	}

	// create a new config loader
	loaderInstance, err := config.NewLoader(opt, *defaultConfig())
	if err != nil {
		log.Fatalf("error loading config: %s", err)
	}

	loader = loaderInstance
}

func C() *Config {
	return loader.C()
}
