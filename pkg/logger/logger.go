package logger

import (
	"context"
	"log/slog"
)

const (
	debugString = "debug"
	infoString  = "info"
	warnString  = "warn"
	errString   = "error"

	Debug LogLevel = debugString
	Info  LogLevel = infoString
	Warn  LogLevel = warnString
	Err   LogLevel = errString
)

type LogLevel string

func (ll LogLevel) String() string {
	return string(ll)
}

type Logger interface {
	Error(ctx context.Context, message string, fields ...slog.Attr)
	Debug(ctx context.Context, message string, fields ...slog.Attr)
	Warn(ctx context.Context, message string, fields ...slog.Attr)
	Info(ctx context.Context, message string, fields ...slog.Attr)
}

func loggingLevelFromStringLevel(lvl string) slog.Level {
	var logLevel slog.Level
	switch level := lvl; level {
	case infoString:
		logLevel = slog.LevelInfo
	case debugString:
		logLevel = slog.LevelDebug
	case warnString:
		logLevel = slog.LevelWarn
	case errString:
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelWarn
	}

	return logLevel
}
