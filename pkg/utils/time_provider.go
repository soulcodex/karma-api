package utils

import (
	"sync"
	"time"
)

type DateTimeProvider interface {
	Now() time.Time
}

type SystemTimeProvider struct{}

func NewSystemTimeProvider() *SystemTimeProvider {
	return &SystemTimeProvider{}
}

func (stp *SystemTimeProvider) Now() time.Time {
	return time.Now()
}

type FixedTimeProvider struct {
	now       time.Time
	timeMutex sync.Mutex
}

func NewFixedTimeProvider() *FixedTimeProvider {
	return &FixedTimeProvider{
		timeMutex: sync.Mutex{},
	}

}

func (ftp *FixedTimeProvider) init() {
	defer ftp.timeMutex.Unlock()
	ftp.timeMutex.Lock()

	if ftp.now.IsZero() {
		ftp.now = time.Now().Round(time.Millisecond)
	}
}

func (ftp *FixedTimeProvider) Now() time.Time {
	ftp.init()
	return ftp.now
}
