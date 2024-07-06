package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App                `yaml:"app"`
		TokenConfiguration `yaml:"token_configuration"`
		HTTP               `yaml:"http"`
		PG                 `yaml:"pg"`
		Redis              `yaml:"redis"`
		JWT                `yaml:"jwt"`
		SMTP               `yaml:"smtp"`
	}

	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	TokenConfiguration struct {
		AccessTokenDuration  int64  `yaml:"access_token_duration"`
		RefreshTokenDuration int64  `yaml:"refresh_token_duration"`
		Issuer               string `yaml:"issuer"`
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

	Redis struct {
		Host     string `yaml:"host" env:"REDIS_HOST"`
		Port     string `yaml:"port" env:"REDIS_PORT"`
		Password string `yaml:"password" env:"REDIS_PASSWORD"`
		DB       int    `yaml:"db" env:"REDIS_DB"`
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
