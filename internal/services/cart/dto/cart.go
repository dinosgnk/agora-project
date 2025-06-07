package dto

type ItemDTO struct {
	ProductId string  `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type AddItemRequestDTO struct {
	UserId   string  `json:"user_id"`
	Item     ItemDTO `json:"item"`
}

type DeleteItemRequestDTO struct {
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
}

type RemoveItemRequestDTO struct {
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
}

type UpdateCartRequestDTO struct {
	UserId	string 	`json:"user_id"`
	Items 	map[string]int `json:"items"`
}

type ClearCartRequest struct {
	UserId 	string `json:"user_id"`
}