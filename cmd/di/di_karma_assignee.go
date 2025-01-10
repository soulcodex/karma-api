package di

import (
	"github.com/soulcodex/karma-api/configs"
	application "github.com/soulcodex/karma-api/internal/karma-assignee/application"
	domain "github.com/soulcodex/karma-api/internal/karma-assignee/domain"
	httpentrypoint "github.com/soulcodex/karma-api/internal/karma-assignee/infrastructure/http"
	persistence "github.com/soulcodex/karma-api/internal/karma-assignee/infrastructure/persistence"
	httpserver "github.com/soulcodex/karma-api/pkg/http-server"
)

const baseKarmaAssigneeRequestsJsonSchemaPath = "karma-assignee/"

type KarmaAssigneeServices struct {
	repository domain.KarmaAssigneeRepository

	// Handlers
	UpsertKarmaAssigneeCommandHandler *application.UpsertKarmaAssigneeCommandHandler
}

func InitKarmaAssigneeServices(c *CommonServices) *KarmaAssigneeServices {
	repository := persistence.NewMySQLKarmaAssigneeRepository(c.DBConnectionPool)

	ks := &KarmaAssigneeServices{
		repository: persistence.NewInMemoryKarmaAssigneeRepository(),
		UpsertKarmaAssigneeCommandHandler: application.NewUpsertKarmaAssigneeCommandHandler(
			repository,
			c.ULIDProvider,
			c.TimeProvider,
		),
	}

	if err := c.CommandBus.RegisterCommand(&application.UpsertKarmaAssigneeCommand{}, ks.UpsertKarmaAssigneeCommandHandler); err != nil {
		panic(err)
	}

	c.RegisterModuleRoutes(registerKarmaAssigneeRoutes())

	return ks
}

func registerKarmaAssigneeRoutes() RouterFunc {
	return func(c *CommonServices, cfg configs.Config) {
		c.Router.Put(
			"/karma-assignees",
			httpentrypoint.HandleUpsertKarmaAssignee(c.CommandBus, c.JsonApiResponseMiddleware),
			httpserver.NewRequestValidatorMiddleware(
				c.JsonApiResponseMiddleware,
				httpserver.RequestValidationJsonSchemaPath(
					cfg.BaseJsonSchemaPath,
					baseKarmaAssigneeRequestsJsonSchemaPath+httpentrypoint.UpsertKarmaAssigneeJsonSchemaPath,
				),
			).Middleware(),
		)
	}
}
