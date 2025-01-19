package karma_assignee_application_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	karma_assignee_application "github.com/soulcodex/karma-api/internal/karma-assignee/application"
	kadomain "github.com/soulcodex/karma-api/internal/karma-assignee/domain"
	karma_assignee_mocks "github.com/soulcodex/karma-api/internal/karma-assignee/domain/mocks"
	"github.com/soulcodex/karma-api/pkg/utils"
)

func TestUpsertKarmaAssigneeCommandHandler(t *testing.T) {
	ulidProvider, timeProvider := utils.NewFixedUlidProvider(), utils.NewFixedTimeProvider()

	tests := []struct {
		name          string
		karmaAssignee func() *kadomain.KarmaAssignee
		expectations  func(context.Context, *kadomain.KarmaAssignee, *karma_assignee_mocks.KarmaAssigneeRepositoryMock)
		errorExpected error
	}{
		{
			name: "it should create a new karma assignee if doesn't exists yet",
			karmaAssignee: func() *kadomain.KarmaAssignee {
				return kadomain.NewKarmaAssigneeFromPrimitives(
					ulidProvider.New().String(),
					"soulcodex",
					"john.doe",
					uint64(1),
					timeProvider.Now(),
				)
			},
			expectations: func(
				ctx context.Context,
				ka *kadomain.KarmaAssignee,
				repo *karma_assignee_mocks.KarmaAssigneeRepositoryMock,
			) {
				notExistErr := kadomain.NewKarmaAssigneeNotExistByUsernameAndAssigner(ka.Username(), ka.Assigner())
				repo.On("FindByUsernameAndAssigner", ctx, ka.Username(), ka.Assigner()).
					Return(nil, notExistErr).
					Once()

				repo.On("Save", ctx, ka).
					Return(nil).
					Once()
			},
			errorExpected: nil,
		},
		{
			name: "it should increase karma counter if exists",
			karmaAssignee: func() *kadomain.KarmaAssignee {
				return kadomain.NewKarmaAssigneeFromPrimitives(
					ulidProvider.New().String(),
					"soulcodex",
					"john.doe",
					uint64(1),
					timeProvider.Now(),
				)
			},
			expectations: func(
				ctx context.Context,
				ka *kadomain.KarmaAssignee,
				repo *karma_assignee_mocks.KarmaAssigneeRepositoryMock,
			) {
				repo.On("FindByUsernameAndAssigner", ctx, ka.Username(), ka.Assigner()).
					Return(ka, nil).
					Once()

				ka = kadomain.NewKarmaAssigneeFromPrimitives(
					ka.Id().String(),
					ka.Username().String(),
					ka.Assigner().String(),
					uint64(2),
					ka.CreatedAt(),
				)
				repo.On("Save", ctx, ka).
					Return(nil).
					Once()
			},
			errorExpected: nil,
		},
		{
			name: "it fails updating a karma assignee if limit has been reached",
			karmaAssignee: func() *kadomain.KarmaAssignee {
				return kadomain.NewKarmaAssigneeFromPrimitives(
					ulidProvider.New().String(),
					"soulcodex",
					"john.doe",
					uint64(5),
					timeProvider.Now(),
				)
			},
			expectations: func(
				ctx context.Context,
				ka *kadomain.KarmaAssignee,
				repo *karma_assignee_mocks.KarmaAssigneeRepositoryMock,
			) {
				repo.On("FindByUsernameAndAssigner", ctx, ka.Username(), ka.Assigner()).
					Return(ka, nil).
					Once()
			},
			errorExpected: &kadomain.KarmaAssigneeLimitReached{},
		},
	}

	for _, scenario := range tests {
		t.Run(scenario.name, func(t *testing.T) {
			ctx, ka := context.Background(), scenario.karmaAssignee()
			repository := karma_assignee_mocks.NewKarmaAssigneeRepositoryMock(t)

			cmd := &karma_assignee_application.UpsertKarmaAssigneeCommand{
				Username: ka.Username().String(),
				Assigner: ka.Assigner().String(),
			}

			handler := karma_assignee_application.NewUpsertKarmaAssigneeCommandHandler(
				repository,
				ulidProvider,
				timeProvider,
			)

			scenario.expectations(ctx, ka, repository)
			err := handler.Handle(ctx, cmd)

			if scenario.errorExpected != nil {
				require.Error(t, err)
				require.ErrorAs(t, err, &scenario.errorExpected)
				return
			}

			require.NoError(t, err)
		})
	}
}
