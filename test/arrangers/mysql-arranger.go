package arrangers

import (
	"context"
	"fmt"
	"sync"

	"github.com/soulcodex/karma-api/pkg/sqldb"
)

type MySQLArranger struct {
	pool           sqldb.ConnectionPool
	migrator       sqldb.DatabaseMigrator
	tablesToIgnore map[string]struct{}
}

func NewMySQLArranger(pool sqldb.ConnectionPool, migrator sqldb.DatabaseMigrator) *MySQLArranger {
	return &MySQLArranger{pool: pool, migrator: migrator, tablesToIgnore: map[string]struct{}{"migrations": {}}}
}

func (msa *MySQLArranger) Arrange(_ context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	msa.mustMigrateUp()
	msa.mustArrangeDB(wg)
}

func (msa *MySQLArranger) mustMigrateUp() {
	_, err := msa.migrator.Up()
	if err != nil {
		panic(err)
	}
}

func (msa *MySQLArranger) mustArrangeDB(wg *sync.WaitGroup) {
	rows, err := msa.pool.Reader().Query("SHOW TABLES")
	if nil != err {
		panic(err)
	}

	defer sqldb.CloseRows(rows)

	var tableName string

	for rows.Next() {
		if scanErr := rows.Scan(&tableName); nil != scanErr {
			panic(scanErr)
		}

		if _, ok := msa.tablesToIgnore[tableName]; ok {
			continue
		}

		truncateSql := fmt.Sprintf("TRUNCATE TABLE %s", tableName)
		if _, truncateErr := msa.pool.Writer().Exec(truncateSql); nil != truncateErr {
			panic(truncateErr)
		}
	}
}
