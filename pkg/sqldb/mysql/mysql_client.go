package xmysql

import (
	"database/sql"
	"fmt"
	"time"
)

const driverName = "mysql"

func NewReader(credentials Credentials, opt ...ClientOptionsFunc) (*sql.DB, error) {
	options := NewDefaultClientOptions(credentials)
	options.apply(opt...)

	sqlAddress := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		options.Credentials.User,
		options.Credentials.Password,
		options.Credentials.Host,
		options.Credentials.Port,
		options.Credentials.Database,
	)

	client, err := sql.Open(driverName, sqlAddress)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(); err != nil {
		return nil, err
	}

	client.SetConnMaxLifetime(time.Duration(options.MaxLifetime) * time.Minute)
	client.SetMaxOpenConns(options.MaxConnections)
	client.SetMaxIdleConns(options.ConnIdle)

	return client, nil
}

func NewWriter(credentials Credentials, opt ...ClientOptionsFunc) (*sql.DB, error) {
	options := NewDefaultClientOptions(credentials)
	options.apply(opt...)

	sqlAddress := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&rejectReadOnly=true",
		options.Credentials.User,
		options.Credentials.Password,
		options.Credentials.Host,
		options.Credentials.Port,
		options.Credentials.Database,
	)

	client, err := sql.Open(driverName, sqlAddress)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(); err != nil {
		return nil, err
	}

	client.SetConnMaxLifetime(time.Duration(options.MaxLifetime) * time.Minute)
	client.SetMaxOpenConns(options.MaxConnections)
	client.SetMaxIdleConns(options.ConnIdle)

	return client, nil
}
