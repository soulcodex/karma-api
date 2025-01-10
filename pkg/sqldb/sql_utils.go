package sqldb

import "database/sql"

func CloseRows(rows *sql.Rows) {
	if rows != nil {
		if err := rows.Close(); err != nil {
			panic(err)
		}
	}
}

func CloseStmt(stmt *sql.Stmt) {
	if stmt != nil {
		if err := stmt.Close(); err != nil {
			panic(err)
		}
	}
}
