package config

import (
	"fmt"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
)

// Option holds configuration options for the loader
type Option struct {
	Prefix       string
	EnvDelimiter string
	Delimiter    string
	YamlFilePath string
}

// CallbackEnv is a function that processes environment variable keys
func (o Option) CallbackEnv(source string) string {
	base := strings.ToLower(strings.TrimPrefix(source, o.Prefix))
	return strings.ReplaceAll(base, o.EnvDelimiter, o.Delimiter)
}

// Loader manages the loading of configuration
type Loader[Config any] struct {
	config *Config
	k      *koanf.Koanf
}

// NewLoader initializes a new ConfigLoader
func NewLoader[Config any](opt Option, defaultConfig Config) (*Loader[Config], error) {
	k := koanf.New(opt.Delimiter)

	// Load default configuration
	if err := k.Load(structs.Provider(defaultConfig, "koanf"), nil); err != nil {
		return nil, fmt.Errorf("error loading default config: %w", err)
	}

	// Load configuration from YAML file
	if err := k.Load(file.Provider(opt.YamlFilePath), yaml.Parser()); err != nil {
		return nil, fmt.Errorf("error loading config from `%s`: %w", opt.YamlFilePath, err)
	}

	// Load environment variables
	if err := k.Load(env.Provider(opt.Prefix, opt.Delimiter, opt.CallbackEnv), nil); err != nil {
		return nil, fmt.Errorf("error loading environment variables: %w", err)
	}

	// Unmarshal loaded config into the provided struct
	var c Config
	if err := k.Unmarshal("", &c); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &Loader[Config]{
		config: &c,
		k:      k,
	}, nil
}

// C returns the loaded configuration
func (cl *Loader[Config]) C() *Config {
	return cl.config
}
