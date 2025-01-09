package logger

import (
	"context"
	"log/slog"
)

type Logger interface {
	Error(ctx context.Context, message string, fields ...slog.Attr)
	Debug(ctx context.Context, message string, fields ...slog.Attr)
	Warn(ctx context.Context, message string, fields ...slog.Attr)
	Info(ctx context.Context, message string, fields ...slog.Attr)
}
