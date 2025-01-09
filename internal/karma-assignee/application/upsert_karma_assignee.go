package karma_assignee_application

import (
	"context"
	"errors"

	kadomain "github.com/soulcodex/karma-api/internal/karma-assignee/domain"
	"github.com/soulcodex/karma-api/pkg/bus"
	"github.com/soulcodex/karma-api/pkg/utils"
)

type UpsertKarmaAssigneeCommand struct {
	Username string
	Assigner string
}

func (uka *UpsertKarmaAssigneeCommand) Id() string {
	return "upsert_karma_assignee_command"
}

type UpsertKarmaAssigneeCommandHandler struct {
	repository   kadomain.KarmaAssigneeRepository
	ulidProvider utils.UlidProvider
	timeProvider utils.DateTimeProvider
}

func NewUpsertKarmaAssigneeCommandHandler(
	r kadomain.KarmaAssigneeRepository,
	up utils.UlidProvider,
	tp utils.DateTimeProvider,
) *UpsertKarmaAssigneeCommandHandler {
	return &UpsertKarmaAssigneeCommandHandler{
		repository:   r,
		ulidProvider: up,
		timeProvider: tp,
	}
}

func (ckh UpsertKarmaAssigneeCommandHandler) Handle(ctx context.Context, command bus.Dto) error {
	cmd, ok := command.(*UpsertKarmaAssigneeCommand)
	if !ok {
		return bus.NewInvalidDto("invalid dto provided")
	}

	id, err := kadomain.NewKarmaAssigneeId(ckh.ulidProvider.New().String())
	if err != nil {
		return nil
	}

	username, err := kadomain.NewUsername(cmd.Username)
	if err != nil {
		return nil
	}

	assigner, err := kadomain.NewUsername(cmd.Assigner)
	if err != nil {
		return nil
	}

	karmaAssignee, findErr := ckh.repository.FindByUsernameAndAssigner(ctx, username, assigner)

	var notExistErr *kadomain.KarmaAssigneeNotExist
	if errors.As(findErr, &notExistErr) {
		ka := kadomain.NewKarmaAssignee(id, username, assigner, ckh.timeProvider.Now())
		return ckh.repository.Save(ctx, ka)
	}

	incrErr := karmaAssignee.IncrementKarma()
	if incrErr != nil {
		return incrErr
	}

	return ckh.repository.Save(ctx, karmaAssignee)
}
