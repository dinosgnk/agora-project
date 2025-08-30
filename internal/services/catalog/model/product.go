package model

type Product struct {
	ProductId   string  `gorm:"primaryKey;column:id"`
	ProductCode string  `gorm:"column:product_code"`
	Name        string  `gorm:"column:name"`
	Category    string  `gorm:"column:category"`
	Description string  `gorm:"column:description"`
	Price       float64 `gorm:"column:price"`
}
