package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	zerolog.Logger
}

// Return new zerologer
func GetLogger() Logger {

	////logging to file
	// file, err := os.OpenFile(
	// 	"myapp.log",
	// 	os.O_APPEND|os.O_CREATE|os.O_WRONLY,
	// 	0664,
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()

	z := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	return Logger{z}
}