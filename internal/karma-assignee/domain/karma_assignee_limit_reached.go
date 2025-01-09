package karma_assignee_domain

import (
	"fmt"

	"github.com/soulcodex/karma-api/pkg/domain"
)

const karmaAssigneeLimitReachedErrorMessage = "karma assignee limit has been reached"

func NewKarmaAssigneeLimitReached(user, assigner Username, counter uint64) *KarmaAssigneeLimitReached {
	return &KarmaAssigneeLimitReached{
		extraItems: map[string]interface{}{
			"user":     user.String(),
			"assigner": assigner.String(),
			"counter":  fmt.Sprintf("%d", counter),
		},
	}
}

type KarmaAssigneeLimitReached struct {
	extraItems map[string]interface{}
	domain.RootDomainError
}

func (uae *KarmaAssigneeLimitReached) Error() string {
	return karmaAssigneeLimitReachedErrorMessage
}

func (uae *KarmaAssigneeLimitReached) ExtraItems() map[string]interface{} {
	return uae.extraItems
}
