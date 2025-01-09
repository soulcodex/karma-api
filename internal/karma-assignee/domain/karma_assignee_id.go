package karma_assignee_domain

import (
	"errors"

	domainvalidation "github.com/soulcodex/karma-api/pkg/domain/validation"
)

var invalidKarmaAssigneeIdProvidedErr = errors.New("invalid karma assignee id provided")

type KarmaAssigneeId string

func NewKarmaAssigneeId(id string) (KarmaAssigneeId, error) {
	karmaAssigneeId := KarmaAssigneeId(id)
	validator := domainvalidation.NewDomainValidator(
		domainvalidation.NotEmpty(),
		domainvalidation.ULIDIdentifier(),
	)
	if err := validator.Validate(id, invalidKarmaAssigneeIdProvidedErr); err != nil {
		return "", err
	}

	return karmaAssigneeId, nil
}

func (kid KarmaAssigneeId) String() string {
	return string(kid)
}
