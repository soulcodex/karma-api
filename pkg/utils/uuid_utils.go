package utils

import "github.com/google/uuid"

type Uuid string

func (u Uuid) String() string {
	return string(u)
}

type UuidProvider interface {
	New() Uuid
}

type RandomUuidProvider struct{}

func NewRandomUuidProvider() *RandomUuidProvider {
	return &RandomUuidProvider{}
}

func (up RandomUuidProvider) New() Uuid {
	return Uuid(uuid.New().String())
}

type FixedUuidProvider struct {
	uuid Uuid
}

func NewFixedUuidProvider() *FixedUuidProvider {
	return &FixedUuidProvider{}
}

func (up *FixedUuidProvider) New() Uuid {
	if up.uuid == "" {
		up.uuid = Uuid(uuid.New().String())
	}

	return up.uuid
}

func GuardUuid(raw string) error {
	_, err := uuid.Parse(raw)
	return err
}

func NewUuid() Uuid {
	return Uuid(uuid.New().String())
}
