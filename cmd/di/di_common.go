package di

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/soulcodex/karma-api/configs"
	"github.com/soulcodex/karma-api/pkg/bus/command"
	"github.com/soulcodex/karma-api/pkg/bus/query"
	distributedsync "github.com/soulcodex/karma-api/pkg/distributed-sync"
	httpserver "github.com/soulcodex/karma-api/pkg/http-server"
	xjsonapi "github.com/soulcodex/karma-api/pkg/json-api"
	xjsonschema "github.com/soulcodex/karma-api/pkg/json-schema"
	"github.com/soulcodex/karma-api/pkg/logger"
	xredis "github.com/soulcodex/karma-api/pkg/redis"
	"github.com/soulcodex/karma-api/pkg/sqldb"
	xmysql "github.com/soulcodex/karma-api/pkg/sqldb/mysql"
	"github.com/soulcodex/karma-api/pkg/utils"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sethvargo/go-envconfig"
)

type CommonServices struct {
	Environment               *configs.Environment
	Config                    configs.Config
	Logger                    logger.Logger
	UUIDProvider              utils.UuidProvider
	ULIDProvider              utils.UlidProvider
	TimeProvider              utils.DateTimeProvider
	JsonApiResponseMiddleware *xjsonapi.JsonApiResponseMiddleware
	JsonSchemaValidator       *xjsonschema.JsonSchemaValidator
	DatabaseMigrator          sqldb.DatabaseMigrator
	DBConnectionPool          sqldb.ConnectionPool
	Router                    *httpserver.Router
	RedisClient               *redis.Client
	MutexService              distributedsync.MutexService
	CommandBus                command.Bus
	QueryBus                  query.Bus

	*RouteRegisterer
}

func InitCommonServices(_ context.Context) *CommonServices {
	configuration := buildConfig()
	environment := configs.MustEnvironment(configuration.AppEnv)
	jsonLogger := buildJSONLogger()
	uuidProvider, ulidProvider := utils.NewRandomUuidProvider(), utils.NewRandomUlidProvider()
	timeProvider := utils.NewSystemTimeProvider()
	httpRouter := buildRouter(configuration, jsonLogger, uuidProvider)
	redisClient := buildRedisClient(configuration)
	mutexService := distributedsync.NewRedisMutexService(redisClient, jsonLogger)
	mysqlPool := buildMySQLConnectionPool(configuration)
	mysqlMigrator := xmysql.NewMysqlDatabaseMigrator(mysqlPool.Writer(), configuration.MigrationsPath, "migrations")

	jsonApiResponseMiddleware := xjsonapi.NewJsonApiResponseMiddleware(jsonLogger)
	jsonSchemaValidator := xjsonschema.NewJsonSchemaValidator(configuration.BaseJsonSchemaPath)
	commandBus := command.InitCommandBus(jsonLogger, mutexService)
	queryBus := query.InitQueryBus(jsonLogger)

	common := &CommonServices{
		Environment:               environment,
		Config:                    configuration,
		Logger:                    jsonLogger,
		UUIDProvider:              uuidProvider,
		ULIDProvider:              ulidProvider,
		TimeProvider:              timeProvider,
		JsonApiResponseMiddleware: jsonApiResponseMiddleware,
		JsonSchemaValidator:       jsonSchemaValidator,
		DBConnectionPool:          mysqlPool,
		DatabaseMigrator:          mysqlMigrator,
		Router:                    httpRouter,
		RedisClient:               redisClient,
		MutexService:              mutexService,
		CommandBus:                commandBus,
		QueryBus:                  queryBus,

		RouteRegisterer: NewRouteRegisterer(httpRouter),
	}

	common.RegisterAllModulesRoutesOnRouter(common)

	return common
}

func buildMySQLConnectionPool(cfg configs.Config) *xmysql.ConnectionPoolMySQL {
	credentials := xmysql.NewCredentials(
		cfg.MysqlUser,
		cfg.MysqlPassword,
		cfg.MysqlHost,
		cfg.MysqlPort,
		cfg.MysqlDatabase,
	)

	reader, err := xmysql.NewReader(credentials)
	if err != nil {
		panic(err)
	}

	writer, err := xmysql.NewWriter(credentials)
	if err != nil {
		panic(err)
	}

	return xmysql.NewMySQLConnectionPool(writer, reader)
}

func buildRedisClient(cfg configs.Config) *redis.Client {
	return xredis.InitRedisClientWithPoolTimeOut(
		cfg.RedisHost,
		cfg.RedisPort,
		cfg.RedisMaxIdleConn,
		cfg.RedisIdleTimeout,
		cfg.RedisIdleTimeout,
	)
}

func RootContext() (context.Context, context.CancelFunc) {
	return context.WithCancel(context.Background())
}

func buildRouter(cfg configs.Config, logger logger.Logger, uuidProvider utils.UuidProvider) *httpserver.Router {
	return httpserver.DefaultRouter(
		cfg.ServerWriteTimeout,
		cfg.ServerReadTimeout,
		httpserver.NewRequestIdentifierMiddleware(uuidProvider).Middleware(),
		httpserver.NewRequestLoggingMiddleware(logger).Middleware,
		httpserver.NewPanicRecoverMiddleware(logger).Middleware(),
	)
}

func buildConfig() configs.Config {
	var cfg configs.Config
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		panic(err)
	}

	return cfg
}

func buildJSONLogger() *logger.JsonStructuredLogger {
	return logger.NewJsonStructuredLogger(logger.Debug)
}
