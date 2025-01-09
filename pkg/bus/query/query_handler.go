package query

import (
	"context"

	"github.com/soulcodex/karma-api/pkg/bus"
)

type QueryHandler interface {
	Handle(ctx context.Context, query bus.Dto) (interface{}, error)
}
