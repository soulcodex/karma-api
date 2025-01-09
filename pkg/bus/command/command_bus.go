package command

import (
	"context"
	"reflect"
	"sync"

	mutex "github.com/soulcodex/karma-api/pkg/distributed-sync"

	"github.com/soulcodex/karma-api/pkg/bus"
	"github.com/soulcodex/karma-api/pkg/logger"
)

type Bus interface {
	RegisterCommand(command bus.Dto, handler CommandHandler) error
	GetHandler(command bus.Dto) (CommandHandler, error)
	Dispatch(ctx context.Context, dto bus.Dto) error
	DispatchAsync(ctx context.Context, dto bus.Dto) error
	ProcessFailed(ctx context.Context)
}

type InMemoryCommandBus struct {
	handlers       map[string]CommandHandler
	lock           sync.Mutex
	logger         logger.Logger
	failedCommands chan *FailedCommand

	mutex mutex.MutexService
}

func InitCommandBus(logger logger.Logger, mutex mutex.MutexService) *InMemoryCommandBus {
	return &InMemoryCommandBus{
		handlers:       make(map[string]CommandHandler),
		lock:           sync.Mutex{},
		logger:         logger,
		failedCommands: make(chan *FailedCommand),

		mutex: mutex,
	}
}

type FailedCommand struct {
	command        bus.Dto
	handler        CommandHandler
	timesProcessed int
}

type AlreadyRegisteredCommand struct {
	message     string
	commandName string
}

func (i AlreadyRegisteredCommand) Error() string {
	return i.message
}

func NewCommandAlreadyRegistered(message string, commandName string) AlreadyRegisteredCommand {
	return AlreadyRegisteredCommand{message: message, commandName: commandName}
}

type NotRegisteredCommandInvoked struct {
	message     string
	commandName string
}

func (i NotRegisteredCommandInvoked) Error() string {
	return i.message
}

func NewCommandNotRegistered(message string, commandName string) NotRegisteredCommandInvoked {
	return NotRegisteredCommandInvoked{message: message, commandName: commandName}
}

func (icb *InMemoryCommandBus) RegisterCommand(command bus.Dto, handler CommandHandler) error {
	icb.lock.Lock()
	defer icb.lock.Unlock()

	commandName, err := icb.commandName(command)
	if err != nil {
		return err
	}

	if _, ok := icb.handlers[*commandName]; ok {
		return NewCommandAlreadyRegistered("command already registered", *commandName)
	}

	icb.handlers[*commandName] = handler

	return nil
}

func (icb *InMemoryCommandBus) GetHandler(command bus.Dto) (CommandHandler, error) {
	commandName, err := icb.commandName(command)
	if err != nil {
		return nil, err
	}
	if handler, ok := icb.handlers[*commandName]; ok {
		return handler, nil
	}

	return nil, NewCommandNotRegistered("command not registered", *commandName)
}

func (icb *InMemoryCommandBus) Dispatch(ctx context.Context, command bus.Dto) error {
	handler, err := icb.GetHandler(command)
	if err != nil {
		return err
	}

	return icb.doHandle(ctx, handler, command)
}

func (icb *InMemoryCommandBus) DispatchAsync(ctx context.Context, command bus.Dto) error {
	commandName, err := icb.commandName(command)
	if err != nil {
		return err
	}

	if handler, ok := icb.handlers[*commandName]; ok {
		go icb.doHandleAsync(ctx, handler, command)

		return nil
	}

	return NewCommandNotRegistered("command not registered", *commandName)
}

func (icb *InMemoryCommandBus) doHandle(ctx context.Context, handler CommandHandler, command bus.Dto) error {
	if bc, ok := command.(bus.BlockOperationCommand); ok {
		operation := func() (interface{}, error) {
			return nil, handler.Handle(ctx, bc)
		}

		_, err := icb.mutex.Mutex(ctx, bc.BlockingKey(), operation)

		return err
	}

	return handler.Handle(ctx, command)
}

func (icb *InMemoryCommandBus) doHandleAsync(ctx context.Context, handler CommandHandler, command bus.Dto) {
	err := icb.doHandle(ctx, handler, command)

	if err != nil {
		icb.failedCommands <- &FailedCommand{
			command:        command,
			handler:        handler,
			timesProcessed: 1,
		}
		icb.logger.Error(ctx, err.Error())
	}
}

func (icb *InMemoryCommandBus) commandName(cmd interface{}) (*string, error) {
	value := reflect.ValueOf(cmd)

	if value.Kind() != reflect.Ptr || !value.IsNil() && value.Elem().Kind() != reflect.Struct {
		return nil, InvalidCommandProvided{"only pointer to commands are allowed"}
	}

	name := value.String()

	return &name, nil
}

func (icb *InMemoryCommandBus) ProcessFailed(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			close(icb.failedCommands)
			icb.logger.Warn(ctx, "exiting safely failed commands consumer...")
			return
		case failedCommand := <-icb.failedCommands:
			if failedCommand.timesProcessed >= 3 {
				continue
			}

			failedCommand.timesProcessed++
			if err := icb.doHandle(ctx, failedCommand.handler, failedCommand.command); err != nil {
				icb.logger.Warn(ctx, err.Error(), logger.ErrValue("previous_error", err))
			}
		}
	}
}

type InvalidCommandProvided struct {
	message string
}

func (i InvalidCommandProvided) Error() string {
	return i.message
}
