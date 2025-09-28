package dto

type Item struct {
	ProductCode string  `json:"product_code"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

type CartResponse struct {
	UserId string `json:"user_id"`
	Items  []Item `json:"items"`
}

type AddItemRequest struct {
	UserId string `json:"user_id"`
	Item   Item   `json:"item"`
}

type RemoveItemRequest struct {
	UserId      string `json:"user_id"`
	ProductCode string `json:"product_code"`
}

type UpdateCartRequest struct {
	UserId string         `json:"user_id"`
	Items  map[string]int `json:"items"`
}

type ClearCartRequest struct {
	UserId string `json:"user_id"`
}
