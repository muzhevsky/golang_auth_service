package internal

import (
	"authorization/utils/errorHandling"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	DatabaseConfig DatabaseConfig `yaml:"database"`
}

type DatabaseConfig struct {
	Connection DatabaseConnectionConfig `yaml:"connection"`
	Tables     []TableConfig            `yaml:"tables"`
}

type DatabaseConnectionConfig struct {
	Hostname     string `yaml:"hostname"`
	Port         string `yaml:"port"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"databaseName"`
}

type TableConfig struct { //TODO make start-up table creation if necessary Tables doesn't exist
	Name   string        `yaml:"name"`
	Fields []FieldConfig `yaml:"fields"`
}

type FieldConfig struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

func NewConfig() *Config {
	reader, err := os.Open("./config/config.yaml")
	errorHandling.LogError(err)

	result := &Config{}
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(result)
	errorHandling.LogError(err)
	return result
}
