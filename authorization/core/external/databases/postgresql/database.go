package postgresql

import (
	"authorization/utils/errorsAndPanics"
	"database/sql"
	_ "github.com/lib/pq"
)

type database struct {
	config         *config
	repositoryCore *RepositoryCore
}

func NewPostgresClient(config *config) *database {
	return &database{config, nil}
}

func (db *database) RepositoryCore() *RepositoryCore {
	return db.repositoryCore
}

func (db *database) Connect() error {
	postgresClient, err := sql.Open("postgres", db.config.ConnectionString())
	errorsAndPanics.HandleError(err)
	errorsAndPanics.HandleError(postgresClient.Ping())
	db.repositoryCore = &RepositoryCore{postgresClient}
	return nil
}

func (db *database) Disconnect() error {
	return nil
}
