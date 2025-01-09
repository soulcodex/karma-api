package utils

import (
	"github.com/oklog/ulid"
	"math/rand"
	"sync"
	"time"
)

type Ulid string

func (u Ulid) String() string {
	return string(u)
}

type UlidProvider interface {
	New() Ulid
}

type RandomUlidProvider struct {
}

func NewRandomUlidProvider() *RandomUlidProvider {
	return &RandomUlidProvider{}
}

func (up RandomUlidProvider) New() Ulid {
	return NewUlid()
}

type FixedUlidProvider struct {
	ulid Ulid
	lock sync.Mutex
}

func NewFixedUlidProvider() *FixedUlidProvider {
	return &FixedUlidProvider{
		lock: sync.Mutex{},
	}
}

func (up *FixedUlidProvider) New() Ulid {
	defer up.lock.Unlock()
	up.lock.Lock()

	if up.ulid == "" {
		up.ulid = NewUlid()
	}

	return up.ulid
}

func GuardUlid(rawUlid string) error {
	_, err := ulid.Parse(rawUlid)
	return err
}

func NewUlid() Ulid {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return Ulid(ulid.MustNew(ulid.Timestamp(t), entropy).String())
}
