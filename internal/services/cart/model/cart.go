package model

type Item struct {
	ProductCode string  `json:"product_code"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

type Cart struct {
	UserId string  `json:"user_id"`
	Items  []*Item `json:"items"`
}
