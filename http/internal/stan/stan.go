package stan

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

type stanClient struct {
	clientID  string
	clusterID string
	natsURL   string
}

// New returns a stan client configuration.
func New(serviceName, clusterID, natsURL string) *stanClient {
	return &stanClient{
		clientID:  fmt.Sprintf("%s-%s", serviceName, uuid.New().String()),
		clusterID: clusterID,
		natsURL:   natsURL,
	}
}

// SendMessage connects to specified client, sends given message and close the stan connection.
func (sc *stanClient) SendMessage(ctx context.Context, channel string, message interface{}) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "stan.SendMessage", ext.SpanKindProducer)
	defer span.Finish()

	client, err := stan.Connect(sc.clusterID, sc.clientID, stan.NatsURL(sc.natsURL))
	if err != nil {
		return err
	}

	injectedMessage, _ := InjectSpanToMap(span, message)
	jsonMessage, err := json.Marshal(injectedMessage)
	if err != nil {
		return err
	}

	span.SetTag("peer.address", sc.natsURL)
	span.SetTag("message_bus.destination", channel)
	span.LogFields(
		log.String("message", "Sending message"),
		log.String("payload", string(jsonMessage)),
	)

	err = client.Publish(channel, jsonMessage)
	if err != nil {
		return err
	}

	err = client.Close()
	if err != nil {
		return err
	}

	return nil
}

// InjectSpanToMap takes a span in context, convert `message` to key value pair,
// injects trace information to a new key, then return it back as NATS payload.
func InjectSpanToMap(span opentracing.Span, message interface{}) (interface{}, error) {
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
