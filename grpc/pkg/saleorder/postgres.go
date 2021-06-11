package saleorder

import (
	"context"
	"encoding/json"

	"github.com/opentracing/opentracing-go"
	"github.com/taogilaaa/trace-sandbox/grpc/internal/database"
	"github.com/taogilaaa/trace-sandbox/grpc/internal/log"
)

type postgresRepository struct {
	db     *database.Store
	logger log.Factory
}

// NewPostgresRepository returns a struct to access data in postgres DB
func NewPostgresRepository(db *database.Store, logger log.Factory) *postgresRepository {
	return &postgresRepository{db, logger}
}

// GetSaleOrder returns a saleorder by given ID.
func (pr *postgresRepository) GetSaleOrder(ctx context.Context, saleOrderID int32) (SaleOrder, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postgres.GetSaleOrder")
	defer span.Finish()

	query := `
		SELECT
			so.id
			, so.email
			, so.payment_method
			, so.order_date
			, COALESCE(
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'name', sop.name
							, 'quantity', sop.quantity
						)
					)
				, '[]') AS products
		FROM
			sale_order so
			LEFT JOIN sale_order_product sop ON so.id = sop.sale_order_id
		WHERE
			so.id = $1
		GROUP BY
			so.id
	`
	var (
		so        SaleOrder
		structstr string
	)

	err := pr.db.GetReplica().QueryRowContext(ctx, query, saleOrderID).Scan(&so.ID, &so.Email, &so.PaymentMethod, &so.OrderDate, &structstr)
	if err != nil {
		pr.logger.For(ctx).WithError(err).Error()
		return so, err
	}

	err = json.Unmarshal([]byte(structstr), &so.Products)
	if err != nil {
		pr.logger.For(ctx).WithError(err).Error()
		return so, err
	}

	return so, nil
}

// GetSaleOrders returns saleorders by given arguments.
func (pr *postgresRepository) GetSaleOrders(ctx context.Context, email string) ([]SaleOrder, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postgres.GetSaleOrders")
	defer span.Finish()

	query := `
		SELECT
			so.id
			, so.email
			, so.payment_method
			, so.order_date
			, COALESCE(
					JSON_AGG(
						JSON_BUILD_OBJECT(
							'name', sop.name
							, 'quantity', sop.quantity
						)
					)
				, '[]') AS products
		FROM
			sale_order so
			LEFT JOIN sale_order_product sop ON so.id = sop.sale_order_id
		WHERE
			(($1 = '') IS NOT FALSE OR so.email = $1)
		GROUP BY
			so.id
	`
	saleOrders := make([]SaleOrder, 0)

	rows, err := pr.db.GetReplica().QueryContext(ctx, query, email)
	if err != nil {
		pr.logger.For(ctx).WithError(err).Error()
		return saleOrders, err
	}
	defer rows.Close()

	for rows.Next() {
		so := SaleOrder{}
		var structstr string

		err := rows.Scan(&so.ID, &so.Email, &so.PaymentMethod, &so.OrderDate, &structstr)
		if err != nil {
			pr.logger.For(ctx).WithError(err).Error()
			return saleOrders, err
		}

		err = json.Unmarshal([]byte(structstr), &so.Products)
		if err != nil {
			pr.logger.For(ctx).WithError(err).Error()
			return saleOrders, err
		}

		saleOrders = append(saleOrders, so)
	}

	if err = rows.Err(); err != nil {
		pr.logger.For(ctx).WithError(err).Error()
		return saleOrders, err
	}

	return saleOrders, nil
}

// CreateSaleOrder is used to create a new SaleOrder.
func (pr *postgresRepository) CreateSaleOrder(ctx context.Context, email, paymentMethod string, products []Product) (int32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "postgres.CreateSaleOrder")
	defer span.Finish()

	var saleOrderID int32
	err := pr.db.GetPrimary().WithTransaction(ctx, func(ctx context.Context, tx database.Tx) error {
		query := `
			INSERT INTO sale_order
				(email, payment_method, order_date)
			VALUES
				($1,    $2,             NOW())
			RETURNING
				id
		`

		err := pr.db.GetPrimary().GetContext(ctx, &saleOrderID, query, email, paymentMethod)
		if err != nil {
			pr.logger.For(ctx).WithError(err).Error()
			return err
		}

		for _, p := range products {
			query = `
				INSERT INTO sale_order_product
					(sale_order_id, name, quantity)
				VALUES
					($1,            $2,   $3)
			`

			_, err := tx.ExecContext(ctx, query, saleOrderID, p.Name, p.Quantity)
			if err != nil {
				pr.logger.For(ctx).WithError(err).Error()
				return err
			}
		}

		return nil
	})
	if err != nil {
		pr.logger.For(ctx).WithError(err).Error()
		return saleOrderID, err
	}

	return saleOrderID, nil
}
