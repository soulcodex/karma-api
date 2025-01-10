package xmysql

type Credentials struct {
	User     string
	Password string
	Host     string
	Port     uint16
	Database string
}

func NewCredentials(user string, password string, host string, port uint16, database string) Credentials {
	return Credentials{
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
	}
}

type ClientOptionsFunc func(co *ClientOptions)

type ClientOptions struct {
	Credentials    Credentials
	MaxConnections int
	ConnIdle       int
	MaxLifetime    int
}

func NewDefaultClientOptions(credentials Credentials) *ClientOptions {
	return &ClientOptions{
		Credentials:    credentials,
		MaxConnections: 20,
		ConnIdle:       50,
		MaxLifetime:    3,
	}
}

func (co *ClientOptions) apply(options ...ClientOptionsFunc) *ClientOptions {
	for _, opt := range options {
		opt(co)
	}
	return co
}

func WithMaxConnections(maxConnections int) ClientOptionsFunc {
	return func(co *ClientOptions) {
		co.MaxConnections = maxConnections
	}
}

func WithConnIdle(connIdle int) ClientOptionsFunc {
	return func(co *ClientOptions) {
		co.ConnIdle = connIdle
	}
}

func WithMaxLifetime(maxLifetime int) ClientOptionsFunc {
	return func(co *ClientOptions) {
		co.MaxLifetime = maxLifetime
	}
}
