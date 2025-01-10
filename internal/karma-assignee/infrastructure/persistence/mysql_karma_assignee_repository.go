package karma_assignee_persistence

import (
	"context"
	"errors"

	kadomain "github.com/soulcodex/karma-api/internal/karma-assignee/domain"
	"github.com/soulcodex/karma-api/pkg/sqldb"
	xmysql "github.com/soulcodex/karma-api/pkg/sqldb/mysql"

	sq "github.com/Masterminds/squirrel"
)

type MySQLKarmaAssigneeRepository struct {
	connPool     sqldb.ConnectionPool
	hydrator     sqldb.HydratorFunc[*kadomain.KarmaAssignee]
	errorHandler *xmysql.ErrorHandler
	tableName    string
	tableFields  []string
}

func NewMySQLKarmaAssigneeRepository(conn sqldb.ConnectionPool) *MySQLKarmaAssigneeRepository {
	return &MySQLKarmaAssigneeRepository{
		connPool: conn,
		hydrator: newMysqlKarmaAssigneeHydrator(),
		errorHandler: xmysql.NewErrorHandler(xmysql.ErrorHandlers{
			xmysql.DuplicatePrimaryKeyErrorCode: uniquenessErrorHandler(),
			xmysql.UniqueViolationErrorCode:     uniquenessErrorHandler(),
		}),
		tableName: "karma_assignee",
		tableFields: []string{
			"id",
			"username",
			"assigner",
			"count",
			"created_at",
		},
	}
}

func (mkr *MySQLKarmaAssigneeRepository) FindByUsernameAndAssigner(
	ctx context.Context,
	username,
	assigner kadomain.Username,
) (*kadomain.KarmaAssignee, error) {
	queryBuilder := sq.Select(mkr.tableFields...).
		From(mkr.tableName).
		Where(sq.Eq{"username": username.String(), "assigner": assigner.String()}).
		OrderBy("created_at DESC").
		Limit(1)

	res, err := queryBuilder.RunWith(mkr.connPool.Reader()).QueryContext(ctx)

	if err != nil {
		return nil, err
	}

	defer sqldb.CloseRows(res)

	if !res.Next() {
		return nil, kadomain.NewKarmaAssigneeNotExistByUsernameAndAssigner(username, assigner)
	}

	karmaAssignee, err := mkr.hydrator(res)
	if err != nil {
		return nil, err
	}

	return karmaAssignee, nil
}

func (mkr *MySQLKarmaAssigneeRepository) Save(ctx context.Context, ka *kadomain.KarmaAssignee) error {
	_, err := sq.Insert(mkr.tableName).
		Columns(mkr.tableFields...).
		Values(
			ka.Id(),
			ka.Username(),
			ka.Assigner(),
			ka.CounterAsNumber(),
			ka.CreatedAt(),
		).
		SuffixExpr(sq.Expr("ON DUPLICATE KEY UPDATE count = ?", ka.CounterAsNumber())).
		RunWith(mkr.connPool.Writer()).
		ExecContext(ctx)

	if err != nil {
		return mkr.errorHandler.Handle(ka, err)
	}

	return nil
}

func uniquenessErrorHandler() func(resource interface{}, err error) error {
	return func(resource interface{}, err error) error {
		ka, match := resource.(*kadomain.KarmaAssignee)
		if !match {
			return errors.New("unhandled resource on repository error handler")
		}

		switch resource.(type) {
		case *kadomain.KarmaAssignee:
			return kadomain.NewKarmaAssigneeAlreadyExist(ka.Id(), ka.Username(), ka.Assigner(), err)
		case kadomain.KarmaAssigneeId:
			return kadomain.NewKarmaAssigneeAlreadyExistWithId(ka.Id(), err)
		default:
			return err
		}
	}
}
