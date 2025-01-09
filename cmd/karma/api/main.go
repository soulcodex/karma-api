package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/soulcodex/karma-api/cmd/di"
)

func main() {
	ctx, cancel := di.RootContext()
	karmaDi, errorsChan := di.InitKarmaDi(ctx), make(chan error)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	defer func() {
		cancel()
	}()

	go func() {
		srvAddress := fmt.Sprintf("%s:%d", karmaDi.Common.Config.ServerHost, karmaDi.Common.Config.ServerPort)
		karmaDi.Common.Logger.Info(ctx, "Starting application...")
		errorsChan <- karmaDi.Common.Router.ListenAndServe(srvAddress)
	}()

	select {
	case <-ctx.Done():
		karmaDi.GracefulShutdown(ctx)
	case <-signals:
		karmaDi.GracefulShutdown(ctx)
	case err := <-errorsChan:
		karmaDi.ErrorShutdown(ctx, cancel, err)
	}
}
