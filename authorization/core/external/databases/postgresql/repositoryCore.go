package postgresql

import "database/sql"

type RepositoryCore struct {
	connection *sql.DB
}
