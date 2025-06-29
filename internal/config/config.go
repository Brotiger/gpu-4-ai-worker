package config

import (
	"github.com/caarlos0/env/v7"
	"log"
)

// TODO: Конфигурация для worker 

type Config struct {
	OllamaImage         string `env:"OLLAMA_IMAGE" envDefault:"ollama/ollama:latest"`
	OllamaContainerName string `env:"OLLAMA_CONTAINER_NAME" envDefault:"ollama"`
	OllamaPort          string `env:"OLLAMA_PORT" envDefault:"11434"`
	VPNConfigPath       string `env:"VPN_CONFIG_PATH" envDefault:"/etc/openvpn/client.conf"`
	VPNUpScript         string `env:"VPN_UP_SCRIPT" envDefault:""`
	GRPCAddr            string `env:"GRPC_ADDR" envDefault:":50051"`
	OllamaDomain        string `env:"OLLAMA_DOMAIN" envDefault:"localhost"`
}

func Load() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("failed to parse env: %v", err)
	}
	return cfg
} 