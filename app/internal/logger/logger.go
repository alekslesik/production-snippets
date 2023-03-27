package logger

import (
	"production-snippets/internal/config"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

// Logger is global logger
var Logger zerolog.Logger

// Return new zerologer
func GetLogger(cfg *config.Config) zerolog.Logger {
	z := zerolog.New(&lumberjack.Logger{
		Filename:   cfg.LoggerSruct.Filename,
		MaxSize:    cfg.LoggerSruct.MaxSize,
		MaxBackups: cfg.LoggerSruct.MaxBackups,
		MaxAge:     cfg.LoggerSruct.MaxAge,
		Compress:   cfg.LoggerSruct.Compress,
	})

	z = z.With().Caller().Timestamp().Logger()
	Logger = z
	
	return z
}
