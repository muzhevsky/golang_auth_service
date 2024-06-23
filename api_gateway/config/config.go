package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP    `yaml:"http"`
		Gateway `yaml:"gateway"`
	}

	HTTP struct {
		Addr string `yaml:"addr"`
		Port string `yaml:"port" env:"API_GATEWAY_PORT"`
	}

	Gateway struct {
		AuthHost string `yaml:"authHost" env:"AUTH_HOST"`
		AuthPort string `yaml:"authPort" env:"AUTH_PORT"`

		ApplicationHost string `yaml:"applicationHost" env:"APPLICATION_HOST"`
		ApplicationPort string `yaml:"applicationPort" env:"APPLICATION_PORT"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	return cfg, nil
}
