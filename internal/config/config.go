package config

import (
	"log"

	"github.com/caarlos0/env/v7"
)

// TODO: Конфигурация для worker

type Config struct {
	OllamaPort   string `env:"OLLAMA_PORT" envDefault:"11434"`
	GRPCAddr     string `env:"GRPC_ADDR" envDefault:":50051"`
	OllamaDomain string `env:"OLLAMA_DOMAIN" envDefault:"localhost"`
}

func Load() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("failed to parse env: %v", err)
	}
	return cfg
}
