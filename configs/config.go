package configs

type Config struct {
	AppEnv             string `env:"APP_ENV"`
	ApplicationName    string `env:"APPLICATION_NAME"`
	ServerHost         string `env:"SERVER_HOST"`
	ServerPort         string `env:"SERVER_PORT"`
	ServerWriteTimeout int    `env:"SERVER_WRITE_TIMEOUT"`
	ServerReadTimeout  int    `env:"SERVER_READ_TIMEOUT"`
}
