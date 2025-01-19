package arrangers

import (
	"context"
	"sync"
)

type Arranger interface {
	Arrange(ctx context.Context, wg *sync.WaitGroup)
}
