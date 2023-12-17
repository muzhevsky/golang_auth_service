package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App    `toml:"app"`
		HTTP   `toml:"http"`
		Logger `toml:"logger"`
		PG     `toml:"pg"`
		JWT    `toml:"jwt"`
		SMTP   `toml:"smtp"`
	}

	App struct {
		Name    string `toml:"name"`
		Version string `toml:"version"`
	}

	HTTP struct {
		Addr string `toml:"addr"`
		Port string `toml:"port"`
	}

	Logger struct {
		Level string `toml:"level"`
	}

	PG struct {
		PoolMax int    `toml:"pool_max"`
		Url     string `toml:"url"`
	}
	JWT struct {
		SigningString string `toml:"signing_string"`
	}
	SMTP struct {
		Username string `toml:"username"`
		Password string `toml:"password"`
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
