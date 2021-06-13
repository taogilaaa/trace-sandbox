package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/taogilaaa/trace-sandbox/http/cmd/server/config"
	"github.com/taogilaaa/trace-sandbox/http/internal/proto/sandbox_sales_v1"
	"github.com/taogilaaa/trace-sandbox/http/internal/stan"
	"github.com/taogilaaa/trace-sandbox/http/pkg/saleorder"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	conn, err := grpc.Dial(
		cfg.GRPCUrl,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal(fmt.Sprintf("error connecting to grpc: %s", err))
	}
	defer conn.Close()

	saleOrderClient := sandbox_sales_v1.NewSaleOrderServiceClient(conn)
	stanClient := stan.New(cfg.AppName, cfg.NATSStreamingCluster, cfg.NATSStreamingUrl)
	httpServer := saleorder.NewHTTPServer(saleOrderClient, stanClient)

	http.HandleFunc("/hello", httpServer.Hello)
	http.HandleFunc("/saleorders", httpServer.SaleOrders)
	http.HandleFunc("/saleorders/", httpServer.SaleOrder)

	http.ListenAndServe(":50042", nil)
}
