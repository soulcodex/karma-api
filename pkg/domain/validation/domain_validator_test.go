package domain_validation_test

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	domain_validation "github.com/soulcodex/karma-api/pkg/domain/validation"
	"github.com/soulcodex/karma-api/pkg/utils"
)

var invalidValueProvidedErr = fmt.Errorf("invalid value provided")

func TestDomainValidatorWithStringConstraints(t *testing.T) {
	uuidProvider, ulidProvider, stringGenerator :=
		utils.NewFixedUuidProvider(),
		utils.NewFixedUlidProvider(),
		utils.NewRandomStringGenerator()

	tests := []struct {
		name         string
		constraints  []domain_validation.DomainValidationRule[string]
		input        string
		expectsError bool
	}{
		{
			name: "it should fail if string is empty",
			constraints: []domain_validation.DomainValidationRule[string]{
				domain_validation.NotEmpty(),
			},
			input:        "",
			expectsError: true,
		},
		{
			name: "it should pass if string isn't empty",
			constraints: []domain_validation.DomainValidationRule[string]{
				domain_validation.NotEmpty(),
			},
			input:        uuidProvider.New().String(),
			expectsError: false,
		},
		{
			name: "it should fail if not match an specific UUID format",
			constraints: []domain_validation.DomainValidationRule[string]{
				domain_validation.NotEmpty(),
				domain_validation.UUIDIdentifier(),
			},
			input:        ulidProvider.New().String(),
			expectsError: true,
		},
		{
			name: "it should pass if match an specific UUID format",
			constraints: []domain_validation.DomainValidationRule[string]{
				domain_validation.NotEmpty(),
				domain_validation.UUIDIdentifier(),
			},
			input:        uuidProvider.New().String(),
			expectsError: false,
		},
		{
			name: "it should fail if not match an specific ULID format",
			constraints: []domain_validation.DomainValidationRule[string]{
				domain_validation.NotEmpty(),
				domain_validation.ULIDIdentifier(),
			},
			input:        uuidProvider.New().String(),
			expectsError: true,
		},
		{
			name: "it should pass if match an specific ULID format",
			constraints: []domain_validation.DomainValidationRule[string]{
				domain_validation.NotEmpty(),
				domain_validation.ULIDIdentifier(),
			},
			input:        ulidProvider.New().String(),
			expectsError: false,
		},
		{
			name: "it should fail if the string is shorten than expected",
			constraints: []domain_validation.DomainValidationRule[string]{
				domain_validation.MinLength(5),
			},
			input:        stringGenerator.Generate(4),
			expectsError: true,
		},
		{
			name: "it should fail if the string is longer than expected",
			constraints: []domain_validation.DomainValidationRule[string]{
				domain_validation.MaxLength(3),
			},
			input:        stringGenerator.Generate(4),
			expectsError: true,
		},
		{
			name: "it should fail if the string doesn't match the pattern",
			constraints: []domain_validation.DomainValidationRule[string]{
				domain_validation.NotEmpty(),
				domain_validation.Regex("^PT-[a-zA-Z]*"),
			},
			input:        stringGenerator.Generate(6),
			expectsError: true,
		},
		{
			name: "it should pass if the string matches the pattern",
			constraints: []domain_validation.DomainValidationRule[string]{
				domain_validation.NotEmpty(),
				domain_validation.Regex("^PT-[a-zA-Z]*"),
			},
			input:        fmt.Sprintf("PT-%s", stringGenerator.Generate(6)),
			expectsError: false,
		},
		{
			name: "it should pass if the string matches the pattern",
			constraints: []domain_validation.DomainValidationRule[string]{
				domain_validation.NotEmpty(),
				domain_validation.Regex("^PT-[a-zA-Z]*"),
			},
			input:        fmt.Sprintf("PT-%s", stringGenerator.Generate(6)),
			expectsError: false,
		},
		{
			name: "it should pass if the string is a valid email",
			constraints: []domain_validation.DomainValidationRule[string]{
				domain_validation.NotEmpty(),
				domain_validation.Email(),
			},
			input:        "john.doe@mail.com",
			expectsError: false,
		},
		{
			name: "it should fail if the string is an invalid email",
			constraints: []domain_validation.DomainValidationRule[string]{
				domain_validation.NotEmpty(),
				domain_validation.Email(),
			},
			input:        "fake.mail",
			expectsError: true,
		},
		{
			name: "it should pass if the string is a valid URL",
			constraints: []domain_validation.DomainValidationRule[string]{
				domain_validation.NotEmpty(),
				domain_validation.URL(),
			},
			input:        fmt.Sprintf("https://%s.nice-domain.eu", strings.ToLower(stringGenerator.Generate(6))),
			expectsError: false,
		},
		{
			name: "it should fail if the string is an invalid URL",
			constraints: []domain_validation.DomainValidationRule[string]{
				domain_validation.NotEmpty(),
				domain_validation.URL(),
			},
			input:        "fake.url",
			expectsError: true,
		},
		{
			name: "it should pass if the string is allowed",
			constraints: []domain_validation.DomainValidationRule[string]{
				domain_validation.NotEmpty(),
				domain_validation.In(map[string]struct{}{
					"accepted": {},
					"rejected": {},
				}),
			},
			input:        "accepted",
			expectsError: false,
		},
		{
			name: "it should fail if the string is not allowed",
			constraints: []domain_validation.DomainValidationRule[string]{
				domain_validation.NotEmpty(),
				domain_validation.In(map[string]struct{}{
					"accepted": {},
					"rejected": {},
				}),
			},
			input:        "delayed",
			expectsError: true,
		},
	}

	for _, scenario := range tests {
		t.Run(scenario.name, func(t *testing.T) {
			validator := domain_validation.NewDomainValidator(scenario.constraints...)
			err := validator.Validate(scenario.input, invalidValueProvidedErr)

			if scenario.expectsError {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestDomainValidatorWithInt64Constraints(t *testing.T) {
	randomNumSeed := rand.NewPCG(uint64(time.Now().UnixMilli()), uint64(time.Now().UnixMilli()))
	numGenerator := rand.New(randomNumSeed)

	tests := []struct {
		name         string
		constraints  []domain_validation.DomainValidationRule[int64]
		input        int64
		expectsError bool
	}{
		{
			name: "it should pass if we provide a valid int64 range",
			constraints: []domain_validation.DomainValidationRule[int64]{
				domain_validation.Int64Range(1, 10),
			},
			input:        int64(1 + numGenerator.IntN(9)),
			expectsError: false,
		},
		{
			name: "it should fail if we provide an invalid int64 out of range",
			constraints: []domain_validation.DomainValidationRule[int64]{
				domain_validation.Int64Range(1, 10),
			},
			input:        int64(11 + numGenerator.IntN(9)),
			expectsError: true,
		},
		{
			name: "it should fail if we provide an insufficient int64 value",
			constraints: []domain_validation.DomainValidationRule[int64]{
				domain_validation.Int64Min(10),
			},
			input:        int64(numGenerator.IntN(9)),
			expectsError: true,
		},
		{
			name: "it should pass if we provide a sufficient int64 value",
			constraints: []domain_validation.DomainValidationRule[int64]{
				domain_validation.Int64Min(10),
			},
			input:        int64(11 + numGenerator.IntN(9)),
			expectsError: false,
		},
		{
			name: "it should fail if we provide a higher value than expected",
			constraints: []domain_validation.DomainValidationRule[int64]{
				domain_validation.Int64Max(10),
			},
			input:        int64(1 + numGenerator.IntN(9)),
			expectsError: true,
		},
		{
			name: "it should pass if we provide a value which not exceed the max limit",
			constraints: []domain_validation.DomainValidationRule[int64]{
				domain_validation.Int64Max(10),
			},
			input:        int64(1 + numGenerator.IntN(9)),
			expectsError: true,
		},
	}

	for _, scenario := range tests {
		t.Run(scenario.name, func(t *testing.T) {
			validator := domain_validation.NewDomainValidator(scenario.constraints...)
			err := validator.Validate(scenario.input, invalidValueProvidedErr)

			if scenario.expectsError {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestDomainValidatorWithFloat64Constraints(t *testing.T) {
	randomNumSeed := rand.NewPCG(uint64(time.Now().UnixMilli()), uint64(time.Now().UnixMilli()))
	numGenerator := rand.New(randomNumSeed)

	tests := []struct {
		name         string
		constraints  []domain_validation.DomainValidationRule[float64]
		input        float64
		expectsError bool
	}{
		{
			name: "it should fail if we provide an invalid float64 out of range",
			constraints: []domain_validation.DomainValidationRule[float64]{
				domain_validation.Float64Range(0.5, 5.5),
			},
			input:        float64(numGenerator.IntN(5)) + 0.5,
			expectsError: false,
		},
		{
			name: "it should pass if we provide a valid float64 on range",
			constraints: []domain_validation.DomainValidationRule[float64]{
				domain_validation.Float64Range(1, 10),
			},
			input:        float64(numGenerator.IntN(9)) + 1.0,
			expectsError: false,
		},
	}

	for _, scenario := range tests {
		t.Run(scenario.name, func(t *testing.T) {
			validator := domain_validation.NewDomainValidator(scenario.constraints...)
			err := validator.Validate(scenario.input, invalidValueProvidedErr)

			if scenario.expectsError {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}
