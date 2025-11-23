package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	zerolog.TimeFieldFormat = time.RFC3339

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
	})
}

func Info(message string, fields ...interface{}) {
	log.Info().Fields(fields).Msg(message)
}

func Error(message string, fields ...interface{}) {
	log.Error().Fields(fields).Msg(message)
}

func Warn(message string, fields ...interface{}) {
	log.Warn().Fields(fields).Msg(message)
}