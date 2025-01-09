package command

import (
	"context"

	"github.com/soulcodex/karma-api/pkg/bus"
)

type CommandHandler interface {
	Handle(ctx context.Context, command bus.Dto) error
}
