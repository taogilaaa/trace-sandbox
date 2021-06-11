package tracing

import (
	"encoding/json"

	"github.com/opentracing/opentracing-go"
)

type tracedMessage struct {
	// TraceID injected by another service, json.Unmarshal is case-insensitive
	// https://golang.org/pkg/encoding/json/#Unmarshal
	TraceID string `json:"Uber-Trace-Id"`
}

// InjectSpanToJSONNats takes a span in context, convert `message` to key value pair,
// injects trace information to a new key, then return it back as NATS payload.
func InjectSpanToJSONNats(span opentracing.Span, message interface{}) (interface{}, error) {
	// Convert interface to key-value pair
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return message, err
	}

	var mappedInterface map[string]interface{}
	err = json.Unmarshal(jsonMessage, &mappedInterface)
	if err != nil {
		return message, err
	}

	// Injects trace context via key value map
	carrier := opentracing.TextMapCarrier{}
	err = span.Tracer().Inject(span.Context(), opentracing.TextMap, carrier)
	if err != nil {
		return message, err
	}

	// Manually inject trace information to payload
	// https://opentracing.io/docs/overview/inject-extract/#inject-pseudocode-example
	for k, v := range carrier {
		mappedInterface[k] = v
	}

	return mappedInterface, nil
}

// ExtractSpanFromJSONNats takes a NATS payload (`MsgProto.Data`) converts `message` to key value pair,
// extracts trace information and return it as a SpanContext.
//
// Return values:
//  - A successful Extract returns a SpanContext instance and a nil error
//  - A bad JSON returns nil SpanContect and the appropriate json error
//  - If there was simply no SpanContext to extract in `carrier`, Extract()
//    returns (nil, opentracing.ErrSpanContextNotFound)
//  - If `format` is unsupported or unrecognized, Extract() returns (nil,
//    opentracing.ErrUnsupportedFormat)
//  - If there are more fundamental problems with the `carrier` object,
//    Extract() may return opentracing.ErrInvalidCarrier,
//    opentracing.ErrSpanContextCorrupted, or implementation-specific
//    errors.
func ExtractSpanFromJSONNats(message []byte) (opentracing.SpanContext, error) {
	var tm tracedMessage
	err := json.Unmarshal(message, &tm)
	if err != nil {
		return nil, err
	}

	// What about other keys such as Zipkin ? should we create multiple keys to support multiple tracers ?
	carrier := opentracing.TextMapCarrier{"Uber-Trace-Id": tm.TraceID}
	return opentracing.GlobalTracer().Extract(opentracing.TextMap, carrier)
}
