package saleorder

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/taogilaaa/trace-sandbox/grpc/internal/log"
)

// Repository provides access to Datastore.
type Repository interface {
	// GetSaleOrder returns a saleorder by given ID.
	GetSaleOrder(ctx context.Context, saleOrderID int32) (SaleOrder, error)
	// GetSaleOrders returns saleorders by given arguments.
	GetSaleOrders(ctx context.Context, email string) ([]SaleOrder, error)
	// CreateSaleOrder is used to create a new SaleOrder.
	CreateSaleOrder(ctx context.Context, email, paymentMethod string, products []Product) (int32, error)
}

type service struct {
	r      Repository
	logger log.Factory
}

// NewService creates a listing service with the necessary dependencies.
func NewService(r Repository, logger log.Factory) *service {
	return &service{r, logger}
}

func (s *service) GetSaleOrder(ctx context.Context, saleOrderID int32) (SaleOrder, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.GetSaleOrder")
	defer span.Finish()

	saleOrder, err := s.r.GetSaleOrder(ctx, saleOrderID)
	if err != nil {
		s.logger.For(ctx).WithError(err).Error()
		return saleOrder, err
	}

	return saleOrder, nil
}

func (s *service) GetSaleOrders(ctx context.Context, email string) ([]SaleOrder, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.GetSaleOrders")
	defer span.Finish()

	saleOrders, err := s.r.GetSaleOrders(ctx, email)
	if err != nil {
		s.logger.For(ctx).WithError(err).Error()
		return saleOrders, err
	}

	return saleOrders, nil
}

func (s *service) CreateSaleOrder(ctx context.Context, email, paymentMethod string, products []Product) (int32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.CreateSaleOrder")
	defer span.Finish()

	if len(products) == 0 {
		s.logger.For(ctx).WithError(ErrEmptyProduct).Error()
		return 0, ErrEmptyProduct
	}

	saleOrderID, err := s.r.CreateSaleOrder(ctx, email, paymentMethod, products)
	if err != nil {
		s.logger.For(ctx).WithError(err).Error()
		return 0, err
	}

	return saleOrderID, nil
}
