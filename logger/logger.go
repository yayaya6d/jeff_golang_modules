package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// logrus.Logger is already thread-safe, we dont need to implement singleton by ourself.
var Log = NewLogger("info")

type Fields map[string]interface{}

// logger is only a wrapper of logrus.logger
type logger struct {
	logrus.Logger
}

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
func NewLogger(level string) Logger {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}

	l := &logger{}
	l.SetOutput(os.Stdout)
	l.SetLevel(logLevel)
	l.SetFormatter(&logrus.JSONFormatter{})

	return l
}

func (l *logger) WithField(key, value string) Logger {
	return &Entry{l.Logger.WithField(key, value)}
}

func (l *logger) WithFields(f Fields) Logger {
	return &Entry{l.Logger.WithFields(logrus.Fields(f))}
}

func (l *logger) SetLoggerLevel(level string) error {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}

	l.SetLevel(logLevel)
	return nil
}

func (l *logger) SetLoggerOutput(output *os.File) {
	l.SetOutput(output)
}
