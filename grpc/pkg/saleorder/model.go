package saleorder

import "time"

type SaleOrder struct {
	ID            int32     `json:"id"`
	Email         string    `json:"email"`
	PaymentMethod string    `json:"payment_method"`
	OrderDate     time.Time `json:"order_date"`
	Products      []Product `json:"products"`
}

type Product struct {
	Name     string `json:"name"`
	Quantity int32  `json:"quantity"`
}
