package configs

import (
	"errors"
	"fmt"
)

const (
	production    EnvironmentName = "production"
	preproduction EnvironmentName = "preproduction"
	test          EnvironmentName = "test"
	local         EnvironmentName = "local"
)

type EnvironmentName string

type Environment struct {
	name EnvironmentName
}

func MustEnvironment(name string) *Environment {
	guardEnvironmentName(name)
	return &Environment{name: EnvironmentName(name)}
}

func environments() map[EnvironmentName]struct{} {
	return map[EnvironmentName]struct{}{
		production:    {},
		preproduction: {},
		test:          {},
		local:         {},
	}
}

func guardEnvironmentName(name string) {
	env := EnvironmentName(name)

	if _, environmentExists := environments()[env]; !environmentExists {
		panic(errors.New(fmt.Sprintf("environment <%s> doesnt exist", name)))
	}
}

func (e *Environment) IsTest() bool {
	return e.name == test
}
