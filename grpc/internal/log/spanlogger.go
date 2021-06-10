package log

import (
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/sirupsen/logrus"
)

type spanLogger struct {
	logger logrus.FieldLogger
	span   opentracing.Span
	fields logrus.Fields
	err    error
}

func (sl spanLogger) WithField(key string, value interface{}) Logger {
	newFields := make(logrus.Fields)
	for key, value := range sl.fields {
		newFields[key] = value
	}
	newFields[key] = value
	sl.fields = newFields
	sl.logger = sl.logger.WithField(key, value)

	return sl
}

func (sl spanLogger) WithFields(fields Fields) Logger {
	newFields := make(logrus.Fields)
	for key, value := range sl.fields {
		newFields[key] = value
	}
	for key, value := range fields {
		newFields[key] = value
	}
	sl.fields = newFields
	sl.logger = sl.logger.WithFields(logrus.Fields(fields))

	return sl
}

func (sl spanLogger) WithError(err error) Logger {
	sl.err = err
	sl.logger = sl.logger.WithError(err)

	return sl
}

func (sl spanLogger) Debug(args ...interface{}) {
	sl.logToSpan("debug", args...)
	sl.logger.Debug(args...)
}

func (sl spanLogger) Info(args ...interface{}) {
	sl.logToSpan("info", args...)
	sl.logger.Info(args...)
}

func (sl spanLogger) Print(args ...interface{}) {
	sl.logToSpan("print", args...)
	sl.logger.Print(args...)
}

func (sl spanLogger) Warn(args ...interface{}) {
	sl.logToSpan("warn", args...)
	sl.logger.Warn(args...)
}

func (sl spanLogger) Warning(args ...interface{}) {
	sl.logToSpan("warning", args...)
	sl.logger.Warning(args...)
}

func (sl spanLogger) Error(args ...interface{}) {
	sl.logToSpan("error", args...)
	sl.logger.Error(args...)
}

func (sl spanLogger) Fatal(args ...interface{}) {
	sl.logToSpan("fatal", args...)
	sl.logger.Fatal(args...)
}

func (sl spanLogger) Panic(args ...interface{}) {
	sl.logToSpan("panic", args...)
	sl.logger.Panic(args...)
}

func (sl spanLogger) logToSpan(level string, args ...interface{}) {
	ef := []log.Field{}
	for key, value := range sl.fields {
		ef = append(ef, log.String(key, fmt.Sprint(value)))
	}

	ef = append(ef, log.String("level", level))
	ef = append(ef, log.String("event", fmt.Sprint(args...)))

	if sl.err != nil {
		ef = append(ef, log.Error(sl.err))
		sl.span.SetTag("error", true)
	}

	sl.span.LogFields(ef...)
}
