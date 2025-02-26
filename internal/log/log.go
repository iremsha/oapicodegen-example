package log

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	sentry "github.com/getsentry/sentry-go"
)

type Attrs map[string]any

type Logger struct {
	logger *slog.Logger
}

type SlogAdapter struct {
	Logger *Logger
}

func (sa *SlogAdapter) Printf(s string, params ...interface{}) {
	sa.Logger.Printf(context.Background(), s, params...)
}

func New() *Logger {
	sentryDsn := os.Getenv("SENTRY_DSN")
	if err := sentry.Init(sentry.ClientOptions{
		AttachStacktrace: true,
		EnableTracing:    true,
		Dsn:              sentryDsn,
	}); err != nil {
		slog.Error("init sentry error", slog.String("error", err.Error()))
	}

	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(time.Now().UTC().Format(time.RFC3339Nano))
			}
			return a
		},
	})
	sentryHook := NewSentryHandler(h, []slog.Level{slog.LevelError})

	logger := slog.New(sentryHook)

	return &Logger{
		logger: logger,
	}
}

func (l *Logger) printLog(ctx context.Context, msg string, level slog.Level, attrs ...Attrs) {
	traceID := ctx.Value("trace_id")
	if traceID == nil {
		traceID = new(string)
	}
	attrsMap := Attrs{}
	if len(attrs) != 0 {
		attrsMap = attrs[0]
	}
	if _, ok := attrsMap["trace_id"]; !ok {
		attrsMap["trace_id"] = traceID
	}
	if _, ok := attrsMap["params"]; !ok {
		attrsMap["params"] = nil
	}

	slogAttrs := make([]slog.Attr, 0, len(attrsMap))
	for key, value := range attrsMap {
		slogAttrs = append(slogAttrs, slog.Any(key, value))
	}

	l.logger.LogAttrs(context.Background(), level, msg, slogAttrs...)
}

func (l *Logger) Error(ctx context.Context, msg string, attrs ...Attrs) {
	l.printLog(ctx, msg, slog.LevelError, attrs...)
}

func (l *Logger) Errorf(ctx context.Context, s string, params ...interface{}) {
	l.Error(ctx, fmt.Sprintf(s, params...))
}

func (l *Logger) Info(ctx context.Context, msg string, attrs ...Attrs) {
	l.printLog(ctx, msg, slog.LevelInfo, attrs...)
}

func (l *Logger) Debug(ctx context.Context, msg string, attrs ...Attrs) {
	l.printLog(ctx, msg, slog.LevelDebug, attrs...)
}

func (l *Logger) Warn(ctx context.Context, msg string, attrs ...Attrs) {
	l.printLog(ctx, msg, slog.LevelWarn, attrs...)
}

func (l *Logger) Printf(ctx context.Context, s string, params ...interface{}) {
	l.Info(ctx, fmt.Sprintf(s, params...))
}

func (l *Logger) Write(d []byte) (n int, err error) {
	s := string(d)
	l.Info(context.Background(), s)
	return len(s), nil
}
