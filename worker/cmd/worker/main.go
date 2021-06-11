package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/nats-io/stan.go"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/taogilaaa/trace-sandbox/worker/cmd/worker/config"
	"github.com/taogilaaa/trace-sandbox/worker/internal/log"
	"github.com/taogilaaa/trace-sandbox/worker/internal/proto/sandbox_sales_v1"
	"github.com/taogilaaa/trace-sandbox/worker/internal/tracing"
	"github.com/taogilaaa/trace-sandbox/worker/pkg/placed"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()
	natsClientId := fmt.Sprintf("%s-%s", cfg.AppName, uuid.New().String())

	tracer, closer := tracing.InitFromEnv(cfg.AppName, cfg.AppVersion)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	baseLogger := logrus.WithFields(logrus.Fields{
		"serviceName":   cfg.AppName,
		"clientVersion": cfg.AppVersion,
		"clientId":      natsClientId,
	})
	baseLogger.Logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "severity",
		},
	})
	logger := log.NewFactory(baseLogger)

	sc, err := stan.Connect(cfg.NATSStreamingCluster, natsClientId, stan.NatsURL(cfg.NATSStreamingUrl),
		stan.Pings(10, 5),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			logger.Bg().WithError(reason).Fatal("Connection Lost")
		}),
	)
	if err != nil {
		logger.Bg().WithError(err).Fatal(fmt.Sprintf("Nats Connection Error: %s", err))
	}
	defer sc.Close()

	logger.Bg().Info("Connected to stan")

	conn, err := grpc.Dial(
		cfg.GRPCUrl,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(tracer)),
		grpc.WithStreamInterceptor(
			otgrpc.OpenTracingStreamClientInterceptor(tracer)),
	)
	if err != nil {
		logger.Bg().WithError(err).Fatal(fmt.Sprintf("error connecting to grpc: %s", err))
	}
	defer conn.Close()

	saleOrderClient := sandbox_sales_v1.NewSaleOrderServiceClient(conn)
	pService := placed.NewService(logger, saleOrderClient)
	pWorker := placed.NewWorker(sc, logger, pService)
	pSubscription, err := pWorker.Run()
	if err != nil {
		sc.Close()
		logger.Bg().WithError(err).Fatal("Subscribe Error")
	}
	defer pSubscription.Close()

	logger.Bg().Info("All Subscriptions ready")

	// Wait for a SIGINT (perhaps triggered by user with CTRL-C)
	// Run cleanup when signal is received
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			baseLogger.Info("\nReceived an interrupt, closing connection...\n\n")
			sc.Close()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}
