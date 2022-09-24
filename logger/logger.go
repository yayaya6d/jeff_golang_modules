package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// logrus.Logger is already thread-safe, we dont need to implement singleton by ourself, just create a instance as singleton object
var Log = NewLogger("info")

type Fields map[string]interface{}

// WrapLogger is only a wrapper of logrus.logger
type WrapLogger struct {
	logrus.Logger
}

// Logger provide user some functions to writing log.
type Logger interface {
	WithField(key, value string) Logger
	WithFields(f Fields) Logger

	Info(args ...interface{})
	Debug(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

// Return a Logger with level to use
func NewLogger(level string) *WrapLogger {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}

	l := &WrapLogger{}
	l.SetOutput(os.Stdout)
	l.SetLevel(logLevel)
	l.SetFormatter(&logrus.JSONFormatter{})

	return l
}

func (l *WrapLogger) WithField(key, value string) Logger {
	return &Entry{l.Logger.WithField(key, value)}
}

func (l *WrapLogger) WithFields(f Fields) Logger {
	return &Entry{l.Logger.WithFields(logrus.Fields(f))}
}

func (l *WrapLogger) SetLoggerLevel(level string) error {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}

	l.SetLevel(logLevel)
	return nil
}
