package karma_assignee_domain

import (
	"sync"
)

const (
	defaultKarmaAssigneeCounter = 1
	maxKarmaAssigneeCount       = 5
)

type KarmaAssigneeCounter struct {
	mutex   sync.Mutex
	counter uint64
}

func (kac *KarmaAssigneeCounter) increment(username, assigner Username) error {
	defer kac.mutex.Unlock()
	kac.mutex.Lock()

	if counter := kac.counter; counter >= maxKarmaAssigneeCount {
		return NewKarmaAssigneeLimitReached(username, assigner, counter)
	}

	kac.counter++
	return nil
}

func NewInitializedKarmaAssigneeCounter() *KarmaAssigneeCounter {
	return &KarmaAssigneeCounter{
		mutex:   sync.Mutex{},
		counter: defaultKarmaAssigneeCounter,
	}
}
