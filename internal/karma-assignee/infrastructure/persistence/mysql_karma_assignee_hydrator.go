package karma_assignee_persistence

import (
	"database/sql"
	"time"

	kadomain "github.com/soulcodex/karma-api/internal/karma-assignee/domain"
	"github.com/soulcodex/karma-api/pkg/sqldb"
)

func newMysqlKarmaAssigneeHydrator() sqldb.HydratorFunc[*kadomain.KarmaAssignee] {
	return func(rows *sql.Rows) (*kadomain.KarmaAssignee, error) {
		var (
			id        string
			username  string
			assigner  string
			count     uint64
			createdAt time.Time
		)

		err := rows.Scan(
			&id, &username, &assigner, &count, &createdAt,
		)
		if err != nil {
			return nil, err
		}

		return kadomain.NewKarmaAssigneeFromPrimitives(
			id,
			username,
			assigner,
			count,
			createdAt,
		), nil
	}
}
