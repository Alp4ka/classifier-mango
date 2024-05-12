package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	// HTTP API.
	HTTPPort  int `env:"HTTP_PORT"`
	RateLimit int `env:"HTTP_RATE_LIMIT" envDefault:"10"`

	// Core service.
	CoreGRPCAddress string `env:"CORE_GRPC_ADDRESS"`

	// Security.
	APIKey string `env:"API_KEY_HASH"`
}

func FromEnv() (*Config, error) {
	cfg := &Config{}

	err := env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("config.FromEnv: %w", err)
	}

	return cfg, nil
}
