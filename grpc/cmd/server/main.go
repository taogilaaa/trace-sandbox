package main

import (
	"fmt"
	"log"
	"net"

	"github.com/taogilaaa/trace-sandbox/grpc/internal/proto/sandbox_sales_v1"
	"github.com/taogilaaa/trace-sandbox/grpc/pkg/saleorder"

	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

const (
	PORT = ":50041"
)

func main() {
	saleorderGrpc := saleorder.NewGRPCServer()
	grpcServer := grpc.NewServer()

	sandbox_sales_v1.RegisterSaleOrderServiceServer(grpcServer, saleorderGrpc)
	healthgrpc.RegisterHealthServer(grpcServer, health.NewServer())

	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal("failed to listen")
	}

	log.Print(fmt.Sprintf("listening at 0.0.0.0%s", PORT))

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve")
	}
}
