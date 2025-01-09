package domain

type Severity int64

const (
	domainErrorSeverity Severity = iota + 1
	criticalErrorSeverity
)

func (sev Severity) Value() int64 {
	return int64(sev)
}

func (sev Severity) IsCritical() bool {
	return sev.Value() == criticalErrorSeverity.Value()
}

func (sev Severity) IsDomainError() bool {
	return sev.Value() == domainErrorSeverity.Value()
}

type RootError interface {
	Error() string
	ExtraItems() map[string]interface{}
	Severity() Severity
	Previous() error
}

type RootCriticalError struct {
	previous error
}

func NewCriticalError() RootCriticalError {
	return RootCriticalError{}
}

func NewCriticalErrorWithPrevious(previous error) RootCriticalError {
	return RootCriticalError{previous: previous}
}

func (rce RootCriticalError) Severity() Severity {
	return criticalErrorSeverity
}

func (rce RootCriticalError) Previous() error {
	return rce.previous
}

func (rce RootCriticalError) Unwrap() error {
	return rce.previous
}

type RootDomainError struct {
	previous error
}

func NewDomainErrorWithPrevious(previous error) RootDomainError {
	return RootDomainError{previous: previous}
}

func (rde RootDomainError) Severity() Severity {
	return domainErrorSeverity
}

func (rde RootDomainError) Previous() error {
	return rde.previous
}
