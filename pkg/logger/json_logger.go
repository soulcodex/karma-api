package logger

import (
	"context"
	"log/slog"
	"os"
)

type JsonStructuredLogger struct {
	logger *slog.Logger
}

func NewJsonStructuredLogger(handler slog.Handler) JsonStructuredLogger {
	if handler != nil {
		return JsonStructuredLogger{logger: slog.New(handler)}
	}

	return JsonStructuredLogger{logger: slog.New(slog.NewJSONHandler(os.Stdout, nil))}
}

func (jsl JsonStructuredLogger) Error(ctx context.Context, message string, fields ...slog.Attr) {
	jsl.logger.LogAttrs(context.Background(), slog.LevelError, message, fields...)
}

func (jsl JsonStructuredLogger) Debug(ctx context.Context, message string, fields ...slog.Attr) {
	jsl.logger.LogAttrs(context.Background(), slog.LevelDebug, message, fields...)
}

func (jsl JsonStructuredLogger) Warn(ctx context.Context, message string, fields ...slog.Attr) {
	jsl.logger.LogAttrs(context.Background(), slog.LevelWarn, message, fields...)
}

func (jsl JsonStructuredLogger) Info(ctx context.Context, message string, fields ...slog.Attr) {
	jsl.logger.LogAttrs(context.Background(), slog.LevelInfo, message, fields...)
}
