package logging

import (
	"context"
	"os"
	"path/filepath"
	"production-snippets/internal/config"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	zerolog.Logger
}

// Return new zerologer
func GetLogger(cfg *config.Config) *Logger {
    var z zerolog.Logger

    path := filepath.Dir(os.Getenv("GOMOD"))

    switch {
    case cfg == nil:
        z = zerolog.New(&lumberjack.Logger{
            Filename:   path + "/default_log.json",
            MaxSize:    100,
            MaxBackups: 3,
            MaxAge:     24,
            Compress:   true,
        })
    default:
        z = zerolog.New(&lumberjack.Logger{
            Filename:   cfg.LoggerSruct.Filename,
            MaxSize:    cfg.LoggerSruct.MaxSize,
            MaxBackups: cfg.LoggerSruct.MaxBackups,
            MaxAge:     cfg.LoggerSruct.MaxAge,
            Compress:   cfg.LoggerSruct.Compress,
        })
        
    }

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	zerolog.TimeFieldFormat = time.DateTime

	z = z.With().Caller().Time("time", time.Now()).Logger()

	return &Logger{z}
}

type ctxLogger struct{}

// ContextWithLogger adds logger to context
func ContextWithLogger(ctx context.Context, l *Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, l)
}

// LoggerFromContext returns logger from context
func LoggerFromContext(ctx context.Context) *Logger {
	if l, ok := ctx.Value(ctxLogger{}).(*Logger); ok {
		return l
	}

	z := GetLogger(&config.Config{})

	return z
}
