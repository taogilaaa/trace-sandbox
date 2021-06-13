package saleorder

type CreateSaleOrder struct {
	Email         string    `json:"email"`
	PaymentMethod string    `json:"payment_method"`
	Products      []Product `json:"products"`
}

type Product struct {
	Name     string `json:"name"`
	Quantity int32  `json:"quantity"`
}
