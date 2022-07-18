package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type ILogger interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warning(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

type Logger struct {
	logger *zerolog.Logger
}

var _ ILogger = (*Logger)(nil)

var lvl zerolog.Level

func New(level string) *Logger {

	switch strings.ToLower(level) {
	case "error":
		lvl = zerolog.ErrorLevel
	case "warn":
		lvl = zerolog.WarnLevel
	case "info":
		lvl = zerolog.InfoLevel
	case "debug":
		lvl = zerolog.DebugLevel
	default:
		lvl = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(1)
	skipFrameCount := 3

	logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return &Logger{
		logger: &logger,
	}
}

func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.build("debug", message, args...)
}

func (l *Logger) Info(message string, args ...interface{}) {
	l.build("info", message, args...)
}

func (l *Logger) Warning(message string, args ...interface{}) {
	l.build("warning", message, args...)
}

func (l *Logger) Error(message interface{}, args ...interface{}) {
	if l.logger.GetLevel() == zerolog.DebugLevel {
		l.Debug(message, args...)
	}

	l.build("error", message, args...)
}

func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.build("debug", message, args...)

	os.Exit(1)
}

func (l *Logger) log(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Info().Msg(message)
	} else {
		l.logger.Info().Msgf(message, args...)
	}
}

func (l *Logger) build(level string, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.log(msg.Error(), args...)
	case string:
		l.log(msg, args...)
	default:
		l.log(fmt.Sprintf("%s message %v has unknown type: %v", level, message, msg), args...)
	}
}
