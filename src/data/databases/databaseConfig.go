package databases

type DatabaseConfig struct {
	connectionString string
	databaseName     string
}

func NewConfig() *DatabaseConfig {
	return &DatabaseConfig{"mongodb://localhost:27017", "englishVocabulary"}
}

func (config *DatabaseConfig) ConnectionString() string {
	return config.connectionString
}

func (config *DatabaseConfig) DatabaseName() string {
	return config.databaseName
}
