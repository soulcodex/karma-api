package domain_validation

import (
	"fmt"
)

func Int64Range(min int64, max int64) DomainValidationRule[int64] {
	return func(value int64) *ValidationError {
		if value >= min && value <= max {
			return nil
		}

		return NewValidationErrorWithMetadata(
			NewValidationMetadata("validation_type", "int64.in_range"),
			NewValidationMetadata("validation_value", fmt.Sprintf("%d", value)),
			NewValidationMetadata("validation_min_range", fmt.Sprintf("%d", min)),
			NewValidationMetadata("validation_max_range", fmt.Sprintf("%d", max)),
		)
	}
}

func Int64Min(min int64) DomainValidationRule[int64] {
	return func(value int64) *ValidationError {
		if value >= min {
			return nil
		}

		return NewValidationErrorWithMetadata(
			NewValidationMetadata("validation_type", "int64.min"),
			NewValidationMetadata("validation_value", fmt.Sprintf("%d", value)),
			NewValidationMetadata("validation_min_range", fmt.Sprintf("%d", min)),
		)
	}
}

func Int64Max(max int64) DomainValidationRule[int64] {
	return func(value int64) *ValidationError {
		if value <= max {
			return nil
		}

		return NewValidationErrorWithMetadata(
			NewValidationMetadata("validation_type", "int64.max"),
			NewValidationMetadata("validation_value", fmt.Sprintf("%d", value)),
			NewValidationMetadata("validation_max_range", fmt.Sprintf("%d", max)),
		)
	}
}

func Float64Range(min float64, max float64) DomainValidationRule[float64] {
	return func(value float64) *ValidationError {
		if value >= min && value <= max {
			return nil
		}

		return NewValidationErrorWithMetadata(
			NewValidationMetadata("validation_type", "float64.in_range"),
			NewValidationMetadata("validation_value", fmt.Sprintf("%f", value)),
			NewValidationMetadata("validation_min_range", fmt.Sprintf("%f", min)),
			NewValidationMetadata("validation_max_range", fmt.Sprintf("%f", max)),
		)
	}
}
