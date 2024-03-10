package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App  `toml:"app"`
		HTTP `toml:"http"`
		PG   `toml:"pg"`
		JWT  `toml:"jwt"`
		SMTP `toml:"smtp"`
	}

	App struct {
		Name    string `toml:"name"`
		Version string `toml:"version"`
	}

	HTTP struct {
		Addr string `toml:"addr"`
		Port string `toml:"port" env:"AUTH_PORT"`
	}

	PG struct {
		PoolMax  string `toml:"pool_max"`
		User     string `toml:"user" env:"POSTGRES_USER"`
		Password string `toml:"password" env:"POSTGRES_PASSWORD"`
		Host     string `toml:"host" env:"POSTGRES_HOST"`
		Port     string `toml:"host" env:"POSTGRES_PORT"`
		Database string `toml:"host" env:"POSTGRES_DATABASE"`
	}
	JWT struct {
		SigningString string `toml:"signing_key" env:"AUTH_JWT_SIGNING_KEY"`
	}
	SMTP struct {
		Username string `toml:"username" env:"AUTH_SMTP_USERNAME"`
		Password string `toml:"password" env:"AUTH_SMTP_PASSWORD"`
		Host     string `toml:"host"`
		Port     string `toml:"port"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./config/config.toml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	return cfg, nil
}
