package logging

import (
	"context"
	"github.com/labstack/echo/v4"
	"log/slog"
	"os"
)

var level = slog.LevelInfo
var additionalKeyValues = map[string]string{}

func ConfigureGlobalLoggerOptions(level slog.Level, metadata map[string]string) {
	level = level
	additionalKeyValues = metadata
}

func IntoContext(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, "logger", l)
}

func FromContext(ctx context.Context) *slog.Logger {
	l := ctx.Value("logger")
	if l == nil {
		return New()
	}
	return l.(*slog.Logger)
}

func New() *slog.Logger {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	for k, v := range additionalKeyValues {
		l = l.With(k, v)
	}
	return l
}

func FromEchoContext(ctx echo.Context) *slog.Logger {
	return FromContext(ctx.Request().Context())
}
