package logger

import "github.com/sirupsen/logrus"

type Entry struct {
	*logrus.Entry
}

func (e *Entry) WithField(key, value string) Logger {
	return &Entry{
		e.Entry.WithField(key, value),
	}
}

func (e *Entry) WithFields(f Fields) Logger {
	return &Entry{
		e.Entry.WithFields(logrus.Fields(f)),
	}
}
