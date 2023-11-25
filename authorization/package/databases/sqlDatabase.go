package databases

import "database/sql"

type ISqlClient interface {
	Connect() error
	Disconnect() error
	GetConnection() *sql.Conn
}
