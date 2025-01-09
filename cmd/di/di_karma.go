package di

import (
	"context"
	"log/slog"
)

type KarmaDi struct {
	Common *CommonServices
}

func InitKarmaDi(ctx context.Context) *KarmaDi {
	return &KarmaDi{
		Common: InitCommonServices(ctx),
	}
}

func (kdi *KarmaDi) ErrorShutdown(ctx context.Context, cancel context.CancelFunc, err error) {
	defer cancel()
	if err == nil {
		return
	}

	kdi.Common.Logger.Error(
		ctx,
		"error starting servers",
		slog.String("service", kdi.Common.Config.ApplicationName),
		slog.String("environment", kdi.Common.Config.AppEnv),
		slog.String("error", err.Error()),
	)
}

func (kdi *KarmaDi) GracefulShutdown(ctx context.Context) {
	kdi.Common.Logger.Info(
		ctx,
		"servers stopped",
		slog.String("service", kdi.Common.Config.ApplicationName),
		slog.String("environment", kdi.Common.Config.AppEnv),
	)
}
