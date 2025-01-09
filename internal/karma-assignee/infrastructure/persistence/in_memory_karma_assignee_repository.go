package karma_assignee_persistence

import (
	"context"
	"sync"

	kadomain "github.com/soulcodex/karma-api/internal/karma-assignee/domain"
)

func karmaAssigneeInMemoryId(username, assigner kadomain.Username) string {
	return assigner.String() + "::" + username.String()
}

type InMemoryKarmaAssigneeRepository struct {
	mutex          sync.RWMutex
	karmaAssignees map[string]*kadomain.KarmaAssignee
}

func NewInMemoryKarmaAssigneeRepository() *InMemoryKarmaAssigneeRepository {
	return &InMemoryKarmaAssigneeRepository{
		mutex:          sync.RWMutex{},
		karmaAssignees: make(map[string]*kadomain.KarmaAssignee),
	}
}

func (kar *InMemoryKarmaAssigneeRepository) FindByUsernameAndAssigner(
	_ context.Context,
	username,
	assigner kadomain.Username,
) (*kadomain.KarmaAssignee, error) {
	defer kar.mutex.RUnlock()
	kar.mutex.RLock()

	id := karmaAssigneeInMemoryId(username, assigner)

	if ka, exists := kar.karmaAssignees[id]; exists {
		return ka, nil
	}

	return nil, kadomain.NewKarmaAssigneeNotExistByUsernameAndAssigner(username, assigner)
}

func (kar *InMemoryKarmaAssigneeRepository) Save(_ context.Context, ka *kadomain.KarmaAssignee) error {
	defer kar.mutex.Unlock()
	kar.mutex.Lock()

	kar.karmaAssignees[karmaAssigneeInMemoryId(ka.Username(), ka.Assigner())] = ka

	return nil
}
