package models

// Order Request
type OrderRequest struct {
	Status       string        `json:"status" validate:"required"`
	Items        []ItemRequest `json:"items" validate:"required"`
	CurrencyUnit string        `json:"currencyUnit" validate:"required"`
}

type ItemRequest struct {
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
	Quantity    uint    `json:"quantity" validate:"required"`
}

// Order Response
type OrderResponse struct {
	OrderID      string         `json:"orderId"`
	Status       string         `json:"status"`
	InvoiceId    string         `json:"invoiceId"`
	Items        []ItemResponse `json:"items"`
	TotalAmount  float32        `json:"totalAmount"`
	CurrencyUnit string         `json:"currencyUnit"`
}

// Items Response
type ItemsResponse struct {
	Items []ItemResponse `json:"items"`
}
type ItemResponse struct {
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    uint    `json:"quantitiy"`
}
