package karma_assignee_domain

import (
	"errors"

	domainvalidation "github.com/soulcodex/karma-api/pkg/domain/validation"
)

const validUsernameRegex = "^[a-zA-Z0-9_\\-.]*$"

var invalidUsernameProvidedErr = errors.New("invalid karma assignee username provided")

type Username string

func NewUsername(u string) (Username, error) {
	username := Username(u)
	validator := domainvalidation.NewDomainValidator(
		domainvalidation.NotEmpty(),
		domainvalidation.Regex(validUsernameRegex),
	)
	if err := validator.Validate(u, invalidUsernameProvidedErr); err != nil {
		return "", err
	}

	return username, nil
}

func (u Username) String() string {
	return string(u)
}
