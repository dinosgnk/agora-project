package model

type Product struct {
	ProductId   int     `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
