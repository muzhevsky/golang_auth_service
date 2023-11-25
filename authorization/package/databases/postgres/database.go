package postgres

import (
	"authorization/package/databases"
	"authorization/utils/errorHandling"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

type postgresClient struct {
	db               *sql.DB
	connectionString *ConnectionString
}

func NewDatabase(connectionString *ConnectionString) databases.ISqlClient {
	result := postgresClient{nil, connectionString}
	return &result
}

func (pg *postgresClient) Connect() error {
	pgDb, err := sql.Open("postgres", pg.connectionString.Build())
	if err != nil {
		return err
	}
	pg.db = pgDb
	err = pgDb.Ping()
	return err
}

func (pg *postgresClient) GetConnection() *sql.Conn {
	connection, err := pg.db.Conn(context.TODO())
	if err != nil {
		errorHandling.LogError(err)
	}
	return connection
}

func (pg *postgresClient) Disconnect() error {
	var err error
	if pg.db != nil {
		err = pg.db.Close()
	}
	return err
}
