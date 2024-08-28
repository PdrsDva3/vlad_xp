package log

import (
	"github.com/rs/zerolog"
	"os"
)

type Logs struct {
	infoLogger  *zerolog.Logger
	errorLogger *zerolog.Logger
}

func (l *Logs) Info(s string) {
	l.infoLogger.Info().Msg(s)
}

func (l *Logs) Error(s string) {
	l.errorLogger.Error().Msg(s)
}

func InitLogger() (*Logs, *os.File, *os.File) {

	loggerInfoFile, err := os.OpenFile("log/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		panic("Error opening info log file")
	}

	loggerErrorFile, err := os.OpenFile("log/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		panic("Error opening error log file")
	}

	infoLogger := zerolog.New(loggerInfoFile).With().Timestamp().Caller().Logger()
	errorLogger := zerolog.New(loggerErrorFile).With().Timestamp().Caller().Logger()

	log := &Logs{
		infoLogger:  &infoLogger,
		errorLogger: &errorLogger,
	}

	return log, loggerInfoFile, loggerErrorFile
}
