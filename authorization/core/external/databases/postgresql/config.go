package postgresql

import (
	"authorization/utils/errorsAndPanics"
	"gopkg.in/yaml.v3"
	"os"
)

type config struct {
	Hostname     string `yaml:"hostname"`
	Port         string `yaml:"port"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"databaseName"`
}

func NewConfig() *config {
	reader, err := os.ReadFile("E:\\Projects\\goLang\\translations\\authorization\\config\\databaseConfig.yaml")
	errorsAndPanics.HandleError(err)

	result := &config{}
	err = yaml.Unmarshal(reader, result)
	errorsAndPanics.HandleError(err)
	return result
}

func (config *config) ConnectionString() string {
	return "postgres://" + config.User + ":" + config.Password + "@" + config.Hostname + ":" + config.Port +
		"/" + config.DatabaseName + "?sslmode=disable"
}
