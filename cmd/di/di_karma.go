package di

import (
	"context"
	"log/slog"
)

type KarmaDi struct {
	Common        *CommonServices
	KarmaAssignee *KarmaAssigneeServices
}

func InitKarmaDi(ctx context.Context) *KarmaDi {
	common := InitCommonServices(ctx)
	karmaAssigneeServices := InitKarmaAssigneeServices(common)

	karmaDI := &KarmaDi{
		Common:        common,
		KarmaAssignee: karmaAssigneeServices,
	}

	common.RegisterAllModulesRoutesOnRouter(common)

	return karmaDI
}

func (kdi *KarmaDi) ErrorShutdown(ctx context.Context, cancel context.CancelFunc, err error) {
	defer cancel()
	if err == nil {
		return
	}

	_ = kdi.Common.Router.Shutdown(ctx)

	kdi.Common.Logger.Error(
		ctx,
		"error starting servers",
		slog.String("service", kdi.Common.Config.ApplicationName),
		slog.String("environment", kdi.Common.Config.AppEnv),
		slog.String("error", err.Error()),
	)
}

func (kdi *KarmaDi) GracefulShutdown(ctx context.Context) {
	_ = kdi.Common.Router.Shutdown(ctx)

	kdi.Common.Logger.Info(
		ctx,
		"servers stopped",
		slog.String("service", kdi.Common.Config.ApplicationName),
		slog.String("environment", kdi.Common.Config.AppEnv),
	)
}
