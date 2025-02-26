package log

import (
	"context"
	"log/slog"
	"slices"

	"github.com/getsentry/sentry-go"
)

type sentryHandler struct {
	slog.Handler
	levels []slog.Level
}

//nolint:revive
func NewSentryHandler(
	handler slog.Handler,
	levels []slog.Level,
) *sentryHandler {
	return &sentryHandler{
		Handler: handler,
		levels:  levels,
	}
}

func (s *sentryHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return s.Handler.Enabled(ctx, level)
}

func (s *sentryHandler) Handle(ctx context.Context, record slog.Record) error {
	if slices.Contains(s.levels, record.Level) {
		if record.Level == slog.LevelError {
			sentry.CaptureMessage(record.Message)
		}
	}
	return s.Handler.Handle(ctx, record)
}

func (s *sentryHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewSentryHandler(s.Handler.WithAttrs(attrs), s.levels)
}

func (s *sentryHandler) WithGroup(group string) slog.Handler {
	return NewSentryHandler(s.Handler.WithGroup(group), s.levels)
}
