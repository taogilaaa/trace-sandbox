package saleorder

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/taogilaaa/trace-sandbox/grpc/internal/log"
	"github.com/taogilaaa/trace-sandbox/grpc/internal/proto/sandbox_sales_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Service provides business logic access.
type Service interface {
	// GetSaleOrder returns a saleorder by given ID.
	GetSaleOrder(ctx context.Context, saleOrderID int32) (SaleOrder, error)
	// GetSaleOrders returns saleorders by given arguments.
	GetSaleOrders(ctx context.Context, email string) ([]SaleOrder, error)
	// CreateSaleOrder is used to create a new SaleOrder.
	CreateSaleOrder(ctx context.Context, email, paymentMethod string, products []Product) (int32, error)
}

type grpcServer struct {
	logger  log.Factory
	service Service
}

// NewServer creates a grpc server with required dependencies.
func NewGRPCServer(logger log.Factory, service Service) *grpcServer {
	return &grpcServer{logger, service}
}

func (gs *grpcServer) GetSaleOrder(ctx context.Context, in *sandbox_sales_v1.GetSaleOrderRequest) (*sandbox_sales_v1.GetSaleOrderResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "grpc.GetSaleOrder")
	defer span.Finish()

	so, err := gs.service.GetSaleOrder(ctx, in.Id)
	if err != nil {
		gs.logger.For(ctx).WithError(err).Error()
		return nil, err
	}

	grpcProducts := make([]*sandbox_sales_v1.Product, 0)
	for _, product := range so.Products {
		grpcProduct := &sandbox_sales_v1.Product{
			Name:     product.Name,
			Quantity: product.Quantity,
		}

		grpcProducts = append(grpcProducts, grpcProduct)
	}

	return &sandbox_sales_v1.GetSaleOrderResponse{
		SaleOrder: &sandbox_sales_v1.SaleOrder{
			Id:            so.ID,
			Email:         so.Email,
			PaymentMethod: so.PaymentMethod,
			OrderDate:     timestamppb.New(so.OrderDate),
			Products:      grpcProducts,
		},
	}, nil
}

func (gs *grpcServer) GetSaleOrders(ctx context.Context, in *sandbox_sales_v1.GetSaleOrdersRequest) (*sandbox_sales_v1.GetSaleOrdersResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "grpc.GetSaleOrders")
	defer span.Finish()

	saleOrders, err := gs.service.GetSaleOrders(ctx, in.Email)
	if err != nil {
		gs.logger.For(ctx).WithError(err).Error()
		return nil, err
	}

	grpcSaleOrders := make([]*sandbox_sales_v1.SaleOrder, 0)
	for _, so := range saleOrders {
		grpcProducts := make([]*sandbox_sales_v1.Product, 0)
		for _, product := range so.Products {
			grpcProduct := &sandbox_sales_v1.Product{
				Name:     product.Name,
				Quantity: product.Quantity,
			}

			grpcProducts = append(grpcProducts, grpcProduct)
		}

		grpcSaleOrders = append(grpcSaleOrders, &sandbox_sales_v1.SaleOrder{
			Id:            so.ID,
			Email:         so.Email,
			PaymentMethod: so.PaymentMethod,
			OrderDate:     timestamppb.New(so.OrderDate),
			Products:      grpcProducts,
		})
	}

	return &sandbox_sales_v1.GetSaleOrdersResponse{
		SaleOrders: grpcSaleOrders,
	}, nil
}

func (gs *grpcServer) CreateSaleOrder(ctx context.Context, in *sandbox_sales_v1.CreateSaleOrderRequest) (*sandbox_sales_v1.CreateSaleOrderResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "grpc.CreateSaleOrder")
	defer span.Finish()

	products := make([]Product, 0)
	for _, p := range in.Products {
		products = append(products, Product{Name: p.Name, Quantity: p.Quantity})
	}

	saleOrderID, err := gs.service.CreateSaleOrder(ctx, in.Email, in.PaymentMethod, products)
	if err != nil {
		gs.logger.For(ctx).WithError(err).Error()
		return nil, err
	}

	return &sandbox_sales_v1.CreateSaleOrderResponse{
		SaleOrder: &sandbox_sales_v1.SaleOrder{
			Id:            saleOrderID,
			Email:         in.Email,
			PaymentMethod: in.PaymentMethod,
			OrderDate:     timestamppb.New(time.Now().UTC()),
			Products:      in.Products,
		},
	}, nil
}
