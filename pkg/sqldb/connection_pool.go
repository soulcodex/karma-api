package sqldb

import "database/sql"

type ConnectionPool interface {
	Writer() *sql.DB
	Reader() *sql.DB
}
