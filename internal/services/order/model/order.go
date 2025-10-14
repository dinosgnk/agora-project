package model

import (
	"time"

	"github.com/dinosgnk/agora-project/internal/services/order/enums"
)

type Order struct {
	ID              string            `gorm:"primaryKey;column:id"`
	UserID          string            `gorm:"column:user_id"`
	Status          enums.OrderStatus `gorm:"column:status"`
	TotalAmount     float64           `gorm:"column:total_amount"`
	ShippingAddress string            `gorm:"column:shipping_address"`
	PaymentMethod   string            `gorm:"column:payment_method"`
	CreatedAt       time.Time         `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time         `gorm:"column:updated_at;autoUpdateTime"`
}

func (Order) TableName() string {
	return "orders.t_order"
}

type OrderedProduct struct {
	ID           string    `gorm:"primaryKey;column:id"`
	OrderID      string    `gorm:"column:order_id"`
	ProductCode  string    `gorm:"column:code"`
	ProductName  string    `gorm:"column:product_name"`
	ProductPrice float64   `gorm:"column:price"`
	Quantity     int       `gorm:"column:quantity"`
	Price        float64   `gorm:"column:price"`
	Subtotal     float64   `gorm:"column:subtotal"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (OrderedProduct) TableName() string {
	return "orders.t_ordered_product"
}

type OrderWithProducts struct {
	Order    Order
	Products []*OrderedProduct
}
