package karma_assignee_domain

import "time"

type KarmaAssignee struct {
	id        KarmaAssigneeId
	username  Username
	assigner  Username
	counter   *KarmaAssigneeCounter
	createdAt time.Time
}

func NewKarmaAssignee(
	id KarmaAssigneeId,
	username Username,
	assigner Username,
	at time.Time,
) *KarmaAssignee {
	return &KarmaAssignee{
		id:        id,
		username:  username,
		assigner:  assigner,
		counter:   NewInitializedKarmaAssigneeCounter(),
		createdAt: at,
	}
}

func NewKarmaAssigneeFromPrimitives(
	id,
	username,
	assigner string,
	count uint64,
	at time.Time,
) *KarmaAssignee {
	return &KarmaAssignee{
		id:        KarmaAssigneeId(id),
		username:  Username(username),
		assigner:  Username(assigner),
		counter:   NewKarmaAssigneeCounterWithCount(count),
		createdAt: at,
	}
}

func (ka *KarmaAssignee) Id() KarmaAssigneeId {
	return ka.id
}

func (ka *KarmaAssignee) Username() Username {
	return ka.username
}

func (ka *KarmaAssignee) Assigner() Username {
	return ka.assigner
}

func (ka *KarmaAssignee) CreatedAt() time.Time {
	return ka.createdAt
}

func (ka *KarmaAssignee) CounterAsNumber() uint64 {
	return ka.counter.counter
}

func (ka *KarmaAssignee) IncrementKarma() error {
	return ka.counter.increment(ka.username, ka.assigner)
}
