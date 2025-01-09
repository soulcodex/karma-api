package karma_assignee_domain

import "context"

type KarmaAssigneeRepository interface {
	FindByUsernameAndAssigner(ctx context.Context, username, assigner Username) (*KarmaAssignee, error)
	Save(ctx context.Context, ka *KarmaAssignee) error
}
