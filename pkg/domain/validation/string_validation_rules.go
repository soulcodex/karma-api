package domain_validation

import (
	"fmt"
	"net/mail"
	"net/url"
	"regexp"
	"strings"

	"github.com/soulcodex/karma-api/pkg/utils"
)

func NotEmpty() DomainValidationRule[string] {
	return func(value string) *ValidationError {
		if value == "" {
			return NewValidationErrorWithMetadata(
				NewValidationMetadata("validation_type", "string.not_empty"),
				NewValidationMetadata("validation_value", value),
			)
		}

		return nil
	}
}

func In(values map[string]struct{}) DomainValidationRule[string] {
	return func(value string) *ValidationError {
		if _, ok := values[value]; ok {
			return nil
		}

		return NewValidationErrorWithMetadata(
			NewValidationMetadata("validation_type", "string.in_values"),
			NewValidationMetadata("validation_value", value),
			NewValidationMetadata(
				"validation_in_values",
				strings.Join(utils.MapStringStructToSlice(values), ","),
			),
		)
	}
}

func MinLength(min int) DomainValidationRule[string] {
	return func(value string) *ValidationError {
		if len(value) < min {
			return NewValidationErrorWithMetadata(
				NewValidationMetadata("validation_type", "string.min_length"),
				NewValidationMetadata("validation_value", value),
				NewValidationMetadata("validation_value_min_length", fmt.Sprintf("%d", min)),
			)
		}

		return nil
	}
}

func MaxLength(max int) DomainValidationRule[string] {
	return func(value string) *ValidationError {
		if len(value) > max {
			return NewValidationErrorWithMetadata(
				NewValidationMetadata("validation_type", "string.max_length"),
				NewValidationMetadata("validation_value", value),
				NewValidationMetadata("validation_value_max_length", fmt.Sprintf("%d", max)),
			)
		}

		return nil
	}
}

func Email() DomainValidationRule[string] {
	return func(value string) *ValidationError {
		if _, emailErr := mail.ParseAddress(value); emailErr != nil {
			return NewValidationErrorWithMetadata(
				NewValidationMetadata("validation_type", "string.email_format"),
				NewValidationMetadata("validation_value", value),
			)
		}

		return nil
	}
}

func URL() DomainValidationRule[string] {
	return func(value string) *ValidationError {
		u, uriErr := url.Parse(value)

		if uriErr == nil && u.Scheme != "" && u.Host != "" {
			return nil
		}

		return NewValidationErrorWithMetadata(
			NewValidationMetadata("validation_type", "string.url_format"),
			NewValidationMetadata("validation_value", value),
		)
	}
}

func Regex(pattern string) DomainValidationRule[string] {
	return func(value string) *ValidationError {
		if matched, matchErr := regexp.MatchString(pattern, value); matchErr != nil || !matched {
			return NewValidationErrorWithMetadata(
				NewValidationMetadata("validation_type", "string.regex_pattern_match"),
				NewValidationMetadata("validation_value", value),
				NewValidationMetadata("validation_value_pattern", pattern),
			)
		}

		return nil
	}
}

func UUIDIdentifier() DomainValidationRule[string] {
	return func(value string) *ValidationError {
		if utils.GuardUuid(value) != nil {
			return NewValidationErrorWithMetadata(
				NewValidationMetadata("validation_type", "string.uuid_identifier_match"),
				NewValidationMetadata("validation_value", value),
			)
		}

		return nil
	}
}

func ULIDIdentifier() DomainValidationRule[string] {
	return func(value string) *ValidationError {
		if utils.GuardUlid(value) != nil {
			return NewValidationErrorWithMetadata(
				NewValidationMetadata("validation_type", "string.ulid_identifier_match"),
				NewValidationMetadata("validation_value", value),
			)
		}

		return nil
	}
}
