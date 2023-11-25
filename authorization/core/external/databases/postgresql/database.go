package postgresql

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type db struct {
	connection *sql.DB
	config     *config
}

func NewDatabase(config *config) *db {
	result := &db{nil, config}
	return result
}

func (db *db) Connect() error {
	postgresClient, err := sql.Open("postgres", db.config.ConnectionString())
	if err != nil {
		return err
	}
	db.connection = postgresClient
	err = postgresClient.Ping()
	return err
}

func (db *db) GetConnection() *sql.DB {
	return db.connection
}

func (db *db) Disconnect() error {
	var err error
	if db.connection != nil {
		err = db.connection.Close()
	}
	return err
}
