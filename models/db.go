package models

// orders Table
type Orders struct {
	ID           uint
	OrderID      string
	OrderStatus  string
	InvoiceID    string
	TotalAmount  float32
	CurrencyUnit string
}

// order_items table
type OrderItems struct {
	ID              uint
	OrderID         uint
	ItemDescription string
	Price           float32
	Quantity        uint
}
