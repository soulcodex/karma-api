package xredis

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	defaultDialTimeout = time.Second * 60
	defaultPoolSize    = 10
	defaultPoolTimeout = time.Second * 60
)

type tracingBuilder func(client *redis.Client) *redis.Client

type TracingOptions struct {
	active         bool
	tracingBuilder tracingBuilder
}

func InitRedisClient(redisHost string, redisPort int, poolSize int, dialTimeout int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", redisHost, redisPort),
		DialTimeout: time.Second * time.Duration(dialTimeout),
		PoolSize:    poolSize,
	})
}

func InitRedisClientWithPoolTimeOut(redisHost string, redisPort int, poolSize int, dialTimeout int, poolTimeOut int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", redisHost, redisPort),
		DialTimeout: time.Second * time.Duration(dialTimeout),
		PoolSize:    poolSize,
		PoolTimeout: time.Second * time.Duration(poolTimeOut),
	})
}

type RedisOptions struct {
	DialTimeout time.Duration
	PoolSize    int
	PoolTimeout time.Duration

	WithTracing TracingOptions
}

type RedisOpt func(opts *RedisOptions)

func WithDialTimeout(dialTimeout time.Duration) RedisOpt {
	return func(opts *RedisOptions) {
		opts.DialTimeout = dialTimeout
	}
}

func WithPoolSize(poolSize int) RedisOpt {
	return func(opts *RedisOptions) {
		opts.PoolSize = poolSize
	}
}
func WithPoolTimeout(poolTimeout time.Duration) RedisOpt {
	return func(opts *RedisOptions) {
		opts.PoolTimeout = poolTimeout
	}
}

func applyOpts(opts []RedisOpt) RedisOptions {
	options := RedisOptions{
		DialTimeout: defaultDialTimeout,
		PoolSize:    defaultPoolSize,
		PoolTimeout: defaultPoolTimeout,
	}
	for _, opt := range opts {
		opt(&options)
	}
	return options
}

func NewRedisClient(host string, port int, opts ...RedisOpt) *redis.Client {
	options := applyOpts(opts)

	redisOptions := &redis.Options{
		Addr:        fmt.Sprintf("%s:%d", host, port),
		DialTimeout: options.DialTimeout,
		PoolSize:    options.PoolSize,
		PoolTimeout: options.PoolTimeout,
	}
	client := redis.NewClient(redisOptions)
	if options.WithTracing.active {
		client = options.WithTracing.tracingBuilder(client)
	}

	return client
}
