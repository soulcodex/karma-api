package arrangers

import (
	"context"
	"sync"

	"github.com/soulcodex/karma-api/cmd/di"
)

type KarmaAPISuiteArranger struct {
	wg            *sync.WaitGroup
	mysqlArranger Arranger
}

func NewKarmaAPISuiteArranger(common *di.CommonServices) *KarmaAPISuiteArranger {
	return &KarmaAPISuiteArranger{
		wg:            &sync.WaitGroup{},
		mysqlArranger: NewMySQLArranger(common.DBConnectionPool, common.DatabaseMigrator),
	}
}

func (ka KarmaAPISuiteArranger) MustArrange(ctx context.Context) {
	ka.wg.Add(1)

	go ka.mysqlArranger.Arrange(ctx, ka.wg)

	ka.wg.Wait()
}
