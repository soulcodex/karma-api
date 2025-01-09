package domain_validation

import (
	"sync"
)

type ValidationErrors struct {
	mutex  sync.RWMutex
	errors []*ValidationError
}

func NewValidationErrors() *ValidationErrors {
	return &ValidationErrors{
		mutex:  sync.RWMutex{},
		errors: make([]*ValidationError, 0),
	}
}

func (v *ValidationErrors) Add(err *ValidationError) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	v.errors = append(v.errors, err)
}

func (v *ValidationErrors) Empty() bool {
	v.mutex.RLock()
	defer v.mutex.RUnlock()

	return len(v.errors) == 0
}

func (v *ValidationErrors) Errors() []*ValidationError {
	v.mutex.RLock()
	defer v.mutex.RUnlock()

	return v.errors
}
