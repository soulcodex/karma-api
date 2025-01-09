package karma_assignee_domain

import (
	"github.com/soulcodex/karma-api/pkg/domain"
)

const karmaAssigneeNotExistErrorMessage = "karma assignee doesn't exist"

func NewKarmaAssigneeNotExistByUsernameAndAssigner(username, assigner Username) *KarmaAssigneeNotExist {
	return &KarmaAssigneeNotExist{
		extraItems: map[string]interface{}{
			"user":     username.String(),
			"assigner": assigner.String(),
		},
	}
}

type KarmaAssigneeNotExist struct {
	extraItems map[string]interface{}
	domain.RootDomainError
}

func (uae *KarmaAssigneeNotExist) Error() string {
	return karmaAssigneeNotExistErrorMessage
}

func (uae *KarmaAssigneeNotExist) ExtraItems() map[string]interface{} {
	return uae.extraItems
}
