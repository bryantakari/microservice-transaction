package inventory

import "time"

type ReservationStatus string

const (
	ItemReserved  ReservationStatus = "CREATED"
	ItemConfirmed ReservationStatus = "CONFIRMED"
	ItemReleased  ReservationStatus = "RELEASED"
	ItemExpired   ReservationStatus = "EXPIRED"
)

type Items struct {
	// Define the structure for the request payload to insert items
	ID               string    `db:"id"`
	Name             string    `db:"item_name"`
	Code             string    `db:"item_code"`
	Description      string    `db:"item_description"`
	Quantity         int       `db:"quantity"`
	ReservedQuantity int       `db:"reserved_quantity"`
	Price            int64     `db:"price"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}

type Transaction struct {
	// Define the structure for the request payload to insert items
	ID        string            `db:"id"`
	ItemId    string            `db:"item_id"`
	Quantity  int               `db:"quantity"`
	Status    ReservationStatus `db:"status"`
	CreatedAt time.Time         `db:"created_at"`
	UpdatedAt time.Time         `db:"updated_at"`
}
