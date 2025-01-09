package logger

import (
	"context"
	"log/slog"
	"os"
)

type JsonStructuredLogger struct {
	logger *slog.Logger
}

func NewJsonStructuredLogger(level LogLevel) *JsonStructuredLogger {
	jsonHandler := slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: loggingLevelFromStringLevel(level.String()),
		},
	)

	return &JsonStructuredLogger{
		logger: slog.New(jsonHandler),
	}
}

func (jsl JsonStructuredLogger) Error(ctx context.Context, message string, fields ...slog.Attr) {
	jsl.logger.LogAttrs(ctx, slog.LevelError, message, fields...)
}

func (jsl JsonStructuredLogger) Debug(ctx context.Context, message string, fields ...slog.Attr) {
	jsl.logger.LogAttrs(ctx, slog.LevelDebug, message, fields...)
}

func (jsl JsonStructuredLogger) Warn(ctx context.Context, message string, fields ...slog.Attr) {
	jsl.logger.LogAttrs(ctx, slog.LevelWarn, message, fields...)
}

func (jsl JsonStructuredLogger) Info(ctx context.Context, message string, fields ...slog.Attr) {
	jsl.logger.LogAttrs(ctx, slog.LevelInfo, message, fields...)
}
