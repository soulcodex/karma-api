package distributed_sync

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"

	"github.com/soulcodex/karma-api/pkg/logger"
	"github.com/soulcodex/karma-api/pkg/utils"
)

const mutexName = "distributed-sync-mutex"

type MutexService interface {
	Mutex(ctx context.Context, key string, fn func() (interface{}, error)) (interface{}, error)
}

type RedisMutexService struct {
	sync   *redsync.Redsync
	logger logger.Logger
}

func NewRedisMutexService(redisClient *redis.Client, logger logger.Logger) *RedisMutexService {
	pool := goredis.NewPool(redisClient)

	return &RedisMutexService{sync: redsync.New(pool), logger: logger}
}

func (rm *RedisMutexService) Mutex(ctx context.Context, key string, fn func() (interface{}, error)) (interface{}, error) {
	mutex := rm.sync.NewMutex(
		mutexName+":"+key,
		redsync.WithExpiry(30*time.Second),
		redsync.WithRetryDelay(25*time.Millisecond),
		redsync.WithTimeoutFactor(0.05),
	)

	if _, acquireErr := utils.RetryFunc(func() (interface{}, error) {
		return nil, rm.acquireLock(ctx, mutex)
	}, 4); acquireErr != nil {
		rm.logger.Error(ctx, "error locking mutex sync", logger.ErrValue("error", acquireErr), slog.String("mutex_key", mutex.Name()))
		return nil, NewErrorLockMutexKey(key, acquireErr)
	}

	result, err := fn()
	if _, releaseErr := utils.RetryFunc(func() (interface{}, error) {
		return nil, rm.releaseLock(ctx, mutex)
	}, 4); releaseErr != nil {
		rm.logger.Error(ctx, "error unlocking mutex sync", logger.ErrValue("error", releaseErr), slog.String("mutex_key", mutex.Name()))
		return nil, NewErrorReleaseLockMutexKey(key, releaseErr)
	}

	return result, err
}

func (rm *RedisMutexService) releaseLock(ctx context.Context, mutex *redsync.Mutex) error {
	if ok, err := mutex.UnlockContext(ctx); !ok || err != nil {
		rm.logger.Warn(ctx, "error unlocking mutex sync - retrying", logger.ErrValue("error", err), slog.String("mutex_key", mutex.Name()))
		if err != nil {
			return err
		}

		return errors.New("redis mutex invalid status when unlocking")
	}

	return nil
}

func (rm *RedisMutexService) acquireLock(ctx context.Context, mutex *redsync.Mutex) error {
	if err := mutex.LockContext(ctx); err != nil {
		rm.logger.Warn(ctx, "error locking mutex sync - retrying", slog.String("error", err.Error()), slog.String("mutex_key", mutex.Name()))
		return err
	}

	return nil
}
