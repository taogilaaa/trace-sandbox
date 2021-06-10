package placed

import (
	"encoding/json"
	"fmt"

	"github.com/nats-io/stan.go"
)

type worker struct {
	sc stan.Conn
}

// NewWorker returns a saleorder placed worker.
func NewWorker(sc stan.Conn) *worker {
	return &worker{sc}
}

// Run connects and execute worker to NATS server.
func (w *worker) Run() (stan.Subscription, error) {
	subscription, err := w.sc.QueueSubscribe(NatsChannel, NatsQueueGroup, func(m *stan.Msg) {
		fmt.Printf("\nIncoming Message %v", string(m.Data))

		so := IncomingMessage{}
		if err := json.Unmarshal(m.Data, &so); err != nil {
			fmt.Printf("%v", err)
			return
		}

		fmt.Println("\n Success")
	}, stan.DurableName(NatsDurableName), stan.MaxInflight(NatsMaxInflight))
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}

	return subscription, nil
}
