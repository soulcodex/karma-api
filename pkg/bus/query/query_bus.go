package query

import (
	"context"

	"sync"

	"github.com/soulcodex/karma-api/pkg/bus"
	"github.com/soulcodex/karma-api/pkg/logger"
)

type Bus interface {
	RegisterQuery(query bus.Dto, handler QueryHandler) error
	Ask(ctx context.Context, dto bus.Dto) (interface{}, error)
}

type InMemoryQueryBus struct {
	handlers map[string]QueryHandler
	lock     sync.Mutex
	logger   logger.Logger
}

func InitQueryBus(logger logger.Logger) *InMemoryQueryBus {
	return &InMemoryQueryBus{
		handlers: make(map[string]QueryHandler, 0),
		lock:     sync.Mutex{},
		logger:   logger,
	}
}

type AlreadyRegisteredQuery struct {
	message   string
	queryName string
}

func (i AlreadyRegisteredQuery) Error() string {
	return i.message
}

func NewQueryAlreadyRegistered(message string, queryName string) AlreadyRegisteredQuery {
	return AlreadyRegisteredQuery{message: message, queryName: queryName}
}

type UnregisteredQuery struct {
	message string
}

func (i UnregisteredQuery) Error() string {
	return i.message
}

func NewQueryNotRegistered(message string, queryName string) AlreadyRegisteredQuery {
	return AlreadyRegisteredQuery{message: message, queryName: queryName}
}

func (bus *InMemoryQueryBus) RegisterQuery(query bus.Dto, handler QueryHandler) error {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	queryName := query.Id()

	if _, ok := bus.handlers[queryName]; ok {
		return NewQueryAlreadyRegistered("query already registered", queryName)
	}

	bus.handlers[queryName] = handler

	return nil
}

func (bus *InMemoryQueryBus) Ask(ctx context.Context, query bus.Dto) (interface{}, error) {
	queryName := query.Id()

	if handler, ok := bus.handlers[queryName]; ok {
		response, err := bus.doAsk(ctx, handler, query)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	return nil, NewQueryNotRegistered("query not registered", queryName)
}

func (bus *InMemoryQueryBus) doAsk(ctx context.Context, handler QueryHandler, query bus.Dto) (interface{}, error) {
	return handler.Handle(ctx, query)
}

type InvalidQueryProvided struct {
	message string
}

func (i InvalidQueryProvided) Error() string {
	return i.message
}
