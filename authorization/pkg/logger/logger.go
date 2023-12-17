package logger

import (
	"github.com/sirupsen/logrus"
	"strings"
)

type ILogger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

type Logger struct {
	logger *logrus.Logger
}

func New(level string) *Logger {
	var lvl logrus.Level

	switch strings.ToLower(level) {
	case "error":
		lvl = logrus.ErrorLevel
	case "warn":
		lvl = logrus.WarnLevel
	case "info":
		lvl = logrus.InfoLevel
	case "debug":
		lvl = logrus.DebugLevel
	default:
		lvl = logrus.InfoLevel
	}

	l := &Logger{
		logger: logrus.New(),
	}

	l.logger.SetLevel(lvl)

	l.logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	return l
}

func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)

}
func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}
func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}
func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}
func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}
