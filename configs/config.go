package configs

type Config struct {
	AppEnv             string `env:"APP_ENV"`
	ApplicationName    string `env:"APPLICATION_NAME"`
	ServerHost         string `env:"SERVER_HOST"`
	ServerPort         int64  `env:"SERVER_PORT"`
	ServerWriteTimeout int    `env:"SERVER_WRITE_TIMEOUT"`
	ServerReadTimeout  int    `env:"SERVER_READ_TIMEOUT"`
	RedisHost          string `env:"REDIS_HOST"`
	RedisPort          int    `env:"REDIS_PORT"`
	RedisMaxIdleConn   int    `env:"REDIS_MAX_IDLE_CONNECTIONS"`
	RedisIdleTimeout   int    `env:"REDIS_IDLE_TIMEOUT"`
	BaseJsonSchemaPath string `env:"BASE_JSON_SCHEMA_PATH"`
}
