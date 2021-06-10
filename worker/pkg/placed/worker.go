package placed

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/taogilaaa/trace-sandbox/worker/internal/log"
)

type worker struct {
	sc     stan.Conn
	logger log.Factory
}

// NewWorker returns a saleorder placed worker.
func NewWorker(sc stan.Conn, logger log.Factory) *worker {
	return &worker{sc, logger}
}

// Run connects and execute worker to NATS server.
func (w *worker) Run() (stan.Subscription, error) {
	subscription, err := w.sc.QueueSubscribe(NatsChannel, NatsQueueGroup, func(m *stan.Msg) {
		span := opentracing.StartSpan(
			fmt.Sprintf("%s.subscription", NatsQueueGroup),
			ext.SpanKindConsumer,
			opentracing.Tag{Key: string(ext.Component), Value: "worker"},
			opentracing.Tag{Key: string(ext.MessageBusDestination), Value: NatsChannel},
		)
		defer span.Finish()
		ctx := opentracing.ContextWithSpan(context.Background(), span)

		requestId := uuid.New().String()
		logger := w.logger.For(ctx).
			WithField("requestId", requestId).
			WithField("sequence", m.Sequence).
			WithField("natsChannel", NatsChannel)

		logger.WithField("data", string(m.Data)).Info("Incoming Message")

		so := IncomingMessage{}
		if err := json.Unmarshal(m.Data, &so); err != nil {
			logger.WithError(err).Error("Json Unmarshal Error")
			return
		}

		logger.Info("Success")
	}, stan.DurableName(NatsDurableName), stan.MaxInflight(NatsMaxInflight))
	if err != nil {
		w.logger.Bg().WithError(err).Fatal("Subscribe Error")
		return nil, err
	}

	w.logger.Bg().Info(fmt.Sprintf("[%s]: Subscription ready", NatsChannel))

	return subscription, nil
}
