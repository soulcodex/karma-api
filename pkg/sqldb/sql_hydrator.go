package sqldb

import "database/sql"

type HydratorFunc[T any] func(rows *sql.Rows) (T, error)
