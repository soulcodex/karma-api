package domain_validation

type DomainValidationRule[T any] func(value T) *ValidationError

type DomainValidator[T any] struct {
	rules []DomainValidationRule[T]
}

func NewDomainValidator[T any](rules ...DomainValidationRule[T]) *DomainValidator[T] {
	return &DomainValidator[T]{
		rules: rules,
	}
}

func (dv *DomainValidator[T]) validate(value T) *ValidationErrors {
	errors := NewValidationErrors()

	for _, rule := range dv.rules {
		if err := rule(value); err != nil {
			errors.Add(err)
		}
	}

	return errors
}

func (dv *DomainValidator[T]) Validate(value T, domainError error) *DomainValidationError {
	errors := dv.validate(value)

	if !errors.Empty() {
		return NewDomainValidationError(errors, domainError)
	}

	return nil
}
