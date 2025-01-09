package domain_validation

import (
	"github.com/soulcodex/karma-api/pkg/domain"
)

const validationErrorDefaultMessage = "validation error has been raised"

type ValidationError struct {
	items map[string]interface{}

	domain.RootDomainError
}

func NewValidationErrorWithMetadata(keyedMetadata ...ValidationMetadata) *ValidationError {
	metadata := make(map[string]interface{}, len(keyedMetadata))
	for _, raw := range keyedMetadata {
		metadata[raw.key] = raw.value
	}

	return &ValidationError{
		items: metadata,
	}
}

func (v *ValidationError) Error() string {
	return validationErrorDefaultMessage
}

func (v *ValidationError) ExtraItems() map[string]interface{} {
	return v.items
}

type ValidationMetadata struct {
	key   string
	value interface{}
}

func NewValidationMetadata(key string, value interface{}) ValidationMetadata {
	return ValidationMetadata{
		key:   key,
		value: value,
	}
}
