package logger

import (
	"log/slog"
)

func ErrValue(key string, err error) slog.Attr {
	return slog.Any(key, err)
}
