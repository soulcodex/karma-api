package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/soulcodex/karma-api/cmd/di"
)

func main() {
	ctx, cancel := di.RootContext()
	karmaDi := di.InitKarmaDi(ctx)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	defer func() {
		cancel()
	}()

	karmaDi.Common.Logger.Info(ctx, "Starting application...")

	select {
	case <-ctx.Done():
		karmaDi.GracefulShutdown(ctx)
	case <-signals:
		karmaDi.GracefulShutdown(ctx)
	}
}
