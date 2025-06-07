package model

type Item struct {
	ProductId string  `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type Cart struct {
	UserId string `json:"user_id"`
	Items  []*Item `json:"items"`
}