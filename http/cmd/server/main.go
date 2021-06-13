package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/cors"
	"github.com/taogilaaa/trace-sandbox/http/cmd/server/config"
	"github.com/taogilaaa/trace-sandbox/http/internal/proto/sandbox_sales_v1"
	"github.com/taogilaaa/trace-sandbox/http/internal/stan"
	"github.com/taogilaaa/trace-sandbox/http/internal/tracing"
	"github.com/taogilaaa/trace-sandbox/http/pkg/saleorder"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	tracer, closer := tracing.InitFromEnv(cfg.AppName, cfg.AppVersion)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	conn, err := grpc.Dial(
		cfg.GRPCUrl,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(tracer)),
		grpc.WithStreamInterceptor(
			otgrpc.OpenTracingStreamClientInterceptor(tracer)),
	)
	if err != nil {
		log.Fatal(fmt.Sprintf("error connecting to grpc: %s", err))
	}
	defer conn.Close()

	saleOrderClient := sandbox_sales_v1.NewSaleOrderServiceClient(conn)
	stanClient := stan.New(cfg.AppName, cfg.NATSStreamingCluster, cfg.NATSStreamingUrl)
	httpServer := saleorder.NewHTTPServer(saleOrderClient, stanClient)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", httpServer.Hello)
	mux.HandleFunc("/saleorders", httpServer.SaleOrders)
	mux.HandleFunc("/saleorders/", httpServer.SaleOrder)

	handler := cors.Default().Handler(mux)

	http.ListenAndServe(":50042", handler)
}
