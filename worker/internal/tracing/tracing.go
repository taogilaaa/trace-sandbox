package tracing

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

// InitFromEnv returns an instance of Jaeger Tracer that read from env
func InitFromEnv(service, version string) (opentracing.Tracer, io.Closer) {
	cfg, err := config.FromEnv()
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}

	cfg.ServiceName = service
	keyTag := fmt.Sprintf("%s.version", service)
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger), config.Tag(keyTag, version), config.MaxTagValueLength(2048))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}

	return tracer, closer
}
