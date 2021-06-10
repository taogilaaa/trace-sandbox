package log

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
)

// Factory is the default logging wrapper that can create
// logger instances either for a given Context or context-less.
type Factory struct {
	logger logrus.FieldLogger
}

// NewFactory creates a new Factory.
func NewFactory(logger logrus.FieldLogger) Factory {
	return Factory{logger: logger}
}

// Bg creates a context-unaware logger.
func (b Factory) Bg() Logger {
	return logger(b)
}

// For returns a context-aware Logger. If the context
// contains an OpenTracing span, all logging calls are also
// echo-ed into the span.
func (b Factory) For(ctx context.Context) Logger {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		var traceID string
		if sc, ok := span.Context().(jaeger.SpanContext); ok {
			traceID = sc.TraceID().String()
		}

		return spanLogger{span: span, logger: b.logger.WithField("trace_id", traceID), fields: logrus.Fields{}}
	}

	return b.Bg()
}
