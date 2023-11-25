package postgresql

import (
	"authorization/utils/errorHandling"
	"gopkg.in/yaml.v3"
	"os"
)

type config struct {
	Connection connectionConfig `yaml:"connection"`
	Tables     []tableConfig    `yaml:"tables"`
}

type connectionConfig struct {
	Hostname     string `yaml:"hostname"`
	Port         string `yaml:"port"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"databaseName"`
}

type tableConfig struct { //TODO make start-up table creation if necessary Tables doesn't exist
	Name   string        `yaml:"name"`
	Fields []fieldConfig `yaml:"fields"`
}

type fieldConfig struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

func NewConfig() *config {
	userDirectory, err := os.UserConfigDir()
	reader, err := os.Open(userDirectory + "\\my_go_apps\\translations\\configs\\databaseConnection.yaml")
	errorHandling.LogError(err)

	result := &config{}
	decoder := yaml.NewDecoder(reader)
	err = decoder.Decode(result)
	errorHandling.LogError(err)
	return result
}

func (config *config) ConnectionString() string {
	cfg := &config.Connection
	return "postgres://" + cfg.User + ":" + cfg.Password + "@" + cfg.Hostname + ":" + cfg.Port + "/" + cfg.DatabaseName + "?sslmode=disable"
}
