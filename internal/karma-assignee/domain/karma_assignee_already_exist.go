package karma_assignee_domain

import (
	"github.com/soulcodex/karma-api/pkg/domain"
)

const karmaAssigneeAlreadyExistErrorMessage = "karma assignee already exist"

func NewKarmaAssigneeAlreadyExist(id KarmaAssigneeId, username, assigner Username, prev error) *KarmaAssigneeAlreadyExist {
	return &KarmaAssigneeAlreadyExist{
		RootDomainError: domain.NewDomainErrorWithPrevious(prev),
		extraItems: map[string]interface{}{
			"id":       id.String(),
			"user":     username.String(),
			"assigner": assigner.String(),
		},
	}
}

func NewKarmaAssigneeAlreadyExistWithId(id KarmaAssigneeId, prev error) *KarmaAssigneeAlreadyExist {
	return &KarmaAssigneeAlreadyExist{
		RootDomainError: domain.NewDomainErrorWithPrevious(prev),
		extraItems: map[string]interface{}{
			"id": id.String(),
		},
	}
}

type KarmaAssigneeAlreadyExist struct {
	extraItems map[string]interface{}
	domain.RootDomainError
}

func (uae *KarmaAssigneeAlreadyExist) Error() string {
	return karmaAssigneeAlreadyExistErrorMessage
}

func (uae *KarmaAssigneeAlreadyExist) ExtraItems() map[string]interface{} {
	return uae.extraItems
}
