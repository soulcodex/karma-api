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

func (ka *KarmaAssignee) IncrementKarma() error {
	return ka.counter.increment(ka.username, ka.assigner)
}
