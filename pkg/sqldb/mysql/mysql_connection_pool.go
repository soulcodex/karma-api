package xmysql

import "database/sql"

type ConnectionPoolMySQL struct {
	writer *sql.DB
	reader *sql.DB
}

func NewMySQLConnectionPool(writer *sql.DB, reader *sql.DB) *ConnectionPoolMySQL {
	guardConnection(writer)
	guardConnection(reader)

	return &ConnectionPoolMySQL{
		writer: writer,
		reader: reader,
	}
}

func NewWithWriterOnly(writer *sql.DB) *ConnectionPoolMySQL {
	guardConnection(writer)

	return &ConnectionPoolMySQL{
		writer: writer,
		reader: nil,
	}
}

func guardConnection(conn *sql.DB) {
	if nil == conn {
		panic(NewInvalidMysqlPoolConfigProvided())
	}
}

func (c *ConnectionPoolMySQL) Writer() *sql.DB {
	return c.writer
}

func (c *ConnectionPoolMySQL) Reader() *sql.DB {
	if nil == c.reader {
		return c.writer
	}

	return c.reader
}
