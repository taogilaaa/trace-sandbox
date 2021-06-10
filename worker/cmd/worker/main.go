package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"github.com/taogilaaa/trace-sandbox/worker/cmd/worker/config"
	"github.com/taogilaaa/trace-sandbox/worker/pkg/placed"
)

func main() {
	cfg := config.Load()
	natsClientId := fmt.Sprintf("%s-%s", cfg.AppName, uuid.New().String())

	sc, err := stan.Connect(cfg.NATSStreamingCluster, natsClientId, stan.NatsURL(cfg.NATSStreamingUrl),
		stan.Pings(10, 5),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection Lost")
		}),
	)
	if err != nil {
		log.Fatal(fmt.Sprintf("Nats Connection Error: %s", err))
	}
	defer sc.Close()

	log.Print("Connected to stan")

	pWorker := placed.NewWorker(sc)
	pSubscription, err := pWorker.Run()
	if err != nil {
		sc.Close()
		log.Fatal("Subscribe Error")
	}
	defer pSubscription.Close()

	log.Print("All Subscriptions ready")

	// Wait for a SIGINT (perhaps triggered by user with CTRL-C)
	// Run cleanup when signal is received
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			log.Print("\nReceived an interrupt, closing connection...\n\n")
			sc.Close()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}
