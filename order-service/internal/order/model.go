package order

import "time"

type OrderStatus string

const (
	OrderCreated   OrderStatus = "CREATED"
	OrderConfirmed OrderStatus = "CONFIRMED"
	OrderCancelled OrderStatus = "CANCELLED"
)

type Order struct {
	ID          string      `db:"id"`
	UserId      string      `db:"user_id"`
	Status      OrderStatus `db:"order_status"`
	TotalAmount int64       `db:"total_order_amount"`
	Currency    string      `db:"currency"`
	CreatedAt   time.Time   `db:"created_at"`
	UpdatedAt   time.Time   `db:"updated_at"`
}

type OrderItem struct {
	ID        string    `db:"id"`
	OrderID   string    `db:"order_id"`
	ProductID string    `db:"product_id"`
	OrderName string    `db:"product_name"`
	Quantity  int       `db:"quantity"`
	UnitPrice int64     `db:"unit_price"`
	CreatedAt time.Time `db:"created_at"`
}
