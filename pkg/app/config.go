package app

import (
	"context"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
)

type Config struct {
	Release      bool
	Port         int
	FetchTimeout int
}

type ConfitaConfig struct {
	Port         int `config:"jackal_port"`
	FetchTimeout int `config:"jackal_fetchtimeout"`
}

func (cfg *ConfitaConfig) ToConfig() *Config {
	return &Config{
		Port:         cfg.Port,
		FetchTimeout: cfg.FetchTimeout,
	}
}

func ConfitaConfigLoader() (*Config, error) {

	cfg := ConfitaConfig{
		Port:         3000,
		FetchTimeout: 2000,
	}

	loader := confita.NewLoader(
		env.NewBackend(),
	)

	err := loader.Load(context.Background(), &cfg)
	if err != nil {
		return nil, err
	}

	return cfg.ToConfig(), nil
}
