package di

import (
	"context"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sethvargo/go-envconfig"

	"github.com/soulcodex/karma-api/configs"
	"github.com/soulcodex/karma-api/pkg/logger"
)

type CommonServices struct {
	Environment *configs.Environment
	Config      configs.Config
	Logger      logger.Logger
}

func InitCommonServices(_ context.Context) *CommonServices {
	configuration := buildConfig()
	environment := configs.MustEnvironment(configuration.AppEnv)
	jsonLogger := buildJSONLogger()

	common := &CommonServices{
		Environment: environment,
		Config:      configuration,
		Logger:      jsonLogger,
	}

	return common
}

func RootContext() (context.Context, context.CancelFunc) {
	return context.WithCancel(context.Background())
}

func buildConfig() configs.Config {
	var cfg configs.Config
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		panic(err)
	}

	return cfg
}

func buildJSONLogger() logger.JsonStructuredLogger {
	return logger.NewJsonStructuredLogger(nil)
}
