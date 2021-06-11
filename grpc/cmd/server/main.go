package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/taogilaaa/trace-sandbox/grpc/cmd/server/config"
	"github.com/taogilaaa/trace-sandbox/grpc/internal/database"
	"github.com/taogilaaa/trace-sandbox/grpc/internal/log"
	"github.com/taogilaaa/trace-sandbox/grpc/internal/proto/sandbox_sales_v1"
	"github.com/taogilaaa/trace-sandbox/grpc/internal/tracing"
	"github.com/taogilaaa/trace-sandbox/grpc/pkg/saleorder"

	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

const (
	PORT = ":50041"
)

func main() {
	cfg := config.Load()
	baseLogger := logrus.WithFields(logrus.Fields{
		"serviceName":   cfg.AppName,
		"clientVersion": cfg.AppVersion,
	})
	baseLogger.Logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "severity",
		},
	})
	baseLogger.Logger.SetReportCaller(true)
	logger := log.NewFactory(baseLogger)

	tracer, closer := tracing.InitFromEnv(cfg.AppName, cfg.AppVersion)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	dataStore := database.New(cfg.Database, database.DriverPostgres)
	defer dataStore.GetPrimary().Close()
	defer dataStore.GetReplica().Close()

	saleorderRepository := saleorder.NewPostgresRepository(dataStore, logger)
	saleorderService := saleorder.NewService(saleorderRepository, logger)
	saleorderGrpc := saleorder.NewGRPCServer(logger, saleorderService)

	grpcTracingOption := otgrpc.IncludingSpans(
		func(parentSpanCtx opentracing.SpanContext, method string, req, resp interface{}) bool {
			// Ignore health check ping
			return !strings.Contains(method, "/grpc.health.v1.Health/")
		},
	)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(tracer, grpcTracingOption, otgrpc.LogPayloads())),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(tracer, grpcTracingOption)),
	)

	sandbox_sales_v1.RegisterSaleOrderServiceServer(grpcServer, saleorderGrpc)
	healthgrpc.RegisterHealthServer(grpcServer, health.NewServer())

	logger.Bg().Info("service registered")

	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		logger.Bg().WithError(err).Fatal("failed to listen")
	}

	logger.Bg().Info(fmt.Sprintf("listening at 0.0.0.0%s", PORT))

	if err := grpcServer.Serve(lis); err != nil {
		logger.Bg().WithError(err).Fatal("failed to serve")
	}
}
