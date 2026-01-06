package order

type CreateOrderRequest struct {
	UserId   string            `json:"user_id"`
	Currency string            `json:"currency"`
	Items    []CreateOrderItem `json:"items"`
}

type CreateOrderItem struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"qty"`
	UnitPrice   int64  `json:"unit_price"`
}
