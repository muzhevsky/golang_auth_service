package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		PG   `yaml:"pg"`
		JWT  `yaml:"jwt"`
		SMTP `yaml:"smtp"`
	}

	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	HTTP struct {
		Addr string `yaml:"addr"`
		Port string `yaml:"port" env:"AUTH_PORT"`
	}

	PG struct {
		PoolMax  string `yaml:"pool_max"`
		User     string `yaml:"user" env:"POSTGRES_USER"`
		Password string `yaml:"password" env:"POSTGRES_PASSWORD"`
		Host     string `yaml:"host" env:"POSTGRES_HOST"`
		Port     string `yaml:"port" env:"POSTGRES_PORT"`
		Database string `yaml:"database" env:"POSTGRES_DATABASE"`
	}
	JWT struct {
		SigningString string `yaml:"signing_key" env:"AUTH_JWT_SIGNING_KEY"`
	}
	SMTP struct {
		Username string `yaml:"username" env:"AUTH_SMTP_USERNAME"`
		Password string `yaml:"password" env:"AUTH_SMTP_PASSWORD"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
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
