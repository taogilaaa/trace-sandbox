package placed

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/taogilaaa/trace-sandbox/worker/internal/log"
	"github.com/taogilaaa/trace-sandbox/worker/internal/proto/sandbox_sales_v1"
)

type service struct {
	logger log.Factory
	sosc   sandbox_sales_v1.SaleOrderServiceClient
}

// NewService creates service with the necessary dependencies.
func NewService(logger log.Factory, grpcClient sandbox_sales_v1.SaleOrderServiceClient) *service {
	return &service{logger, grpcClient}
}

func (s *service) CreateSaleOrder(ctx context.Context, saleOrder IncomingMessage) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.CreateSaleOrder", ext.SpanKindRPCClient)
	defer span.Finish()

	if len(saleOrder.Products) == 0 {
		s.logger.For(ctx).WithError(ErrEmptyProduct).Error("CreateSaleOrder Error")
		return ErrEmptyProduct
	}

	products := make([]*sandbox_sales_v1.Product, 0)
	for _, p := range saleOrder.Products {
		products = append(products, &sandbox_sales_v1.Product{Name: p.Name, Quantity: p.Quantity})
	}

	payload := &sandbox_sales_v1.CreateSaleOrderRequest{
		Email:         saleOrder.Email,
		PaymentMethod: saleOrder.PaymentMethod,
		Products:      products,
	}

	_, err := s.sosc.CreateSaleOrder(ctx, payload)
	if err != nil {
		s.logger.For(ctx).WithError(err).Error("CreateSaleOrder Error")
		return err
	}

	return nil
}
