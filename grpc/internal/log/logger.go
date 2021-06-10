package log

import "github.com/sirupsen/logrus"

// Logger is a simplified abstraction of the logrus.FieldLogger
type Logger interface {
	// Adds a field to the log entry, note that it doesn't log until you call
	// Debug, Print, Info, Warn, Error, Fatal or Panic. It only creates a log entry.
	// If you want multiple fields, use `WithFields`.
	WithField(key string, value interface{}) Logger
	// Adds a struct of fields to the log entry. All it does is call `WithField` for
	// each `Field`.
	WithFields(fields Fields) Logger
	// Add an error as single field to the log entry.  All it does is call
	// `WithError` for the given `error`.
	WithError(err error) Logger

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

type Fields logrus.Fields

// logger delegates all calls to the underlying logrus.FieldLogger
type logger struct {
	logger logrus.FieldLogger
}

func (l logger) WithField(key string, value interface{}) Logger {
	return logger{logger: l.logger.WithField(key, value)}
}

func (l logger) WithFields(fields Fields) Logger {
	return logger{logger: l.logger.WithFields(logrus.Fields(fields))}
}

func (l logger) WithError(err error) Logger {
	return logger{logger: l.logger.WithError(err)}
}

func (l logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l logger) Print(args ...interface{}) {
	l.logger.Print(args...)
}

func (l logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l logger) Warning(args ...interface{}) {
	l.logger.Warning(args...)
}

func (l logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l logger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}
