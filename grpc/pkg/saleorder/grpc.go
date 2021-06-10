package saleorder

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/taogilaaa/trace-sandbox/grpc/internal/proto/sandbox_sales_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type grpcServer struct {
}

// NewServer creates a grpc server with required dependencies.
func NewGRPCServer() *grpcServer {
	return &grpcServer{}
}

func (gs *grpcServer) GetSaleOrder(ctx context.Context, in *sandbox_sales_v1.GetSaleOrderRequest) (*sandbox_sales_v1.GetSaleOrderResponse, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "grpc.GetSaleOrder")
	defer span.Finish()

	return &sandbox_sales_v1.GetSaleOrderResponse{
		SaleOrder: &sandbox_sales_v1.SaleOrder{
			Id:            1,
			Email:         "",
			PaymentMethod: "",
			OrderDate:     timestamppb.New(time.Now().UTC()),
			Products:      []*sandbox_sales_v1.Product{},
		},
	}, nil
}

func (gs *grpcServer) GetSaleOrders(ctx context.Context, in *sandbox_sales_v1.GetSaleOrdersRequest) (*sandbox_sales_v1.GetSaleOrdersResponse, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "grpc.GetSaleOrders")
	defer span.Finish()

	return &sandbox_sales_v1.GetSaleOrdersResponse{
		SaleOrders: []*sandbox_sales_v1.SaleOrder{},
	}, nil
}

func (gs *grpcServer) CreateSaleOrder(ctx context.Context, in *sandbox_sales_v1.CreateSaleOrderRequest) (*sandbox_sales_v1.CreateSaleOrderResponse, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "grpc.CreateSaleOrder")
	defer span.Finish()

	return &sandbox_sales_v1.CreateSaleOrderResponse{
		SaleOrder: &sandbox_sales_v1.SaleOrder{
			Id:            1,
			Email:         "",
			PaymentMethod: "",
			OrderDate:     timestamppb.New(time.Now().UTC()),
			Products:      []*sandbox_sales_v1.Product{},
		},
	}, nil
}
