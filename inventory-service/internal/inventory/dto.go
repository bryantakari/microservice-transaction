package inventory

type ReservationResult string

const (
	ReservationReserved          ReservationResult = "RESERVED"
	ReservationInsufficientStock ReservationResult = "INSUFFICIENT_STOCK"
	ReservationInvalidItem       ReservationResult = "INVALID_ITEM"
	ReservationDuplicateOrder    ReservationResult = "DUPLICATE_ORDER"
)

type InsertItemsRequest struct {
	// Define the structure for the request payload to insert items
	Items []ItemsDTO `json:"items"`
}

type ItemsDTO struct {
	ID          string `json:"id"`
	Name        string `json:"item_name"`
	Description string `json:"item_description"`
	Price       int64  `json:"price"`
	Code        string `json:"item_code"`
	Quantity    int    `json:"quantity"`
}

type ListItemsQuery struct {
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
	SortBy  string `json:"sort_by"`
	OrderBy string `json:"order_by"` // asc | desc
}

type ReservationResponse struct {
	// Define the structure for the request payload to insert items
	ReservationStatusResult ReservationResult `json:"reservation_status"`
	ReservationID           string            `json:"reservation_id"`
}

type ReserveItemsRequest struct {
	// Define the structure for the request payload to reserve items
	Order_ID string            `json:"order_id"`
	Items    []ReserveItemsDto `json:"items"`
}

type ReserveItemsDto struct {
	ItemId   string `json:"item_id"`
	Quantity int    `json:"quantity"`
}
