package domain_validation

import (
	"github.com/soulcodex/karma-api/pkg/domain"
)

const rootDomainErrorsDetailsKey = "details"
const rootDomainValidationErrorMessage = "domain validation error occurred"

type DomainValidationError struct {
	domain.RootDomainError

	items map[string]interface{}
}

func NewDomainValidationError(errors *ValidationErrors, domainErr error) *DomainValidationError {
	return &DomainValidationError{
		RootDomainError: domain.NewDomainErrorWithPrevious(domainErr),
		items:           buildDomainValidationErrorExtraItemsWithErrorDetails(errors),
	}
}

func (rve *DomainValidationError) Error() string {
	return rootDomainValidationErrorMessage
}

func (rve *DomainValidationError) ExtraItems() map[string]interface{} {
	return rve.items
}

func (rve *DomainValidationError) ErrorDetails() []map[string]interface{} {
	errors := make([]map[string]interface{}, 0)

	if rve.items == nil {
		return errors
	}

	validationErrors, detailsKeyExists := rve.items[rootDomainErrorsDetailsKey]
	if !detailsKeyExists {
		return errors
	}

	errors = append(errors, validationErrors.([]map[string]interface{})...)

	return errors
}

func buildDomainValidationErrorExtraItemsWithErrorDetails(errors *ValidationErrors) map[string]interface{} {
	details := map[string]interface{}{
		rootDomainErrorsDetailsKey: make([]map[string]interface{}, 0),
	}

	for _, e := range errors.Errors() {
		details[rootDomainErrorsDetailsKey] = append(
			details[rootDomainErrorsDetailsKey].([]map[string]interface{}),
			e.ExtraItems(),
		)
	}

	return details
}
