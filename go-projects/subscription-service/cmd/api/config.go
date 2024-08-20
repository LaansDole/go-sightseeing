package main

import (
	"context"
	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
)

type Config struct {
	Host string `json:"host" yaml:"host" default:"localhost"`
	Port int    `json:"port" yaml:"port" default:"8080"`
}

var configFile string

func loadConfig(configFile string) (ctx context.Context, cfg Config, err error) {
	var contents []byte
	contents, err = os.ReadFile(configFile)
	if err != nil {
		return
	}

	// Unmarshal
	err = yaml.Unmarshal(contents, &cfg)
	if err != nil {
		return
	}
	cfg, err = loadPort(cfg)

	// Set defaults and port
	if err = defaults.Set(&cfg); err != nil {
		return
	}

	// Setup context
	ctx = context.Background()
	return
}

func loadPort(cfg Config) (Config, error) {
	var err error
	if port := os.Getenv("PORT"); port != "" {
		cfg.Port, err = strconv.Atoi(port)
	}
	return cfg, err
}
