package inventory

type InsertItemsRequest struct {
	// Define the structure for the request payload to insert items
	Items []ItemsDTO `json:"items"`
}

type ItemsDTO struct {
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
