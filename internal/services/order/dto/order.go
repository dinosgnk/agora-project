package dto

import (
	"time"

	"github.com/dinosgnk/agora-project/internal/services/order/enums"
)

type OrderedProduct struct {
	ProductCode string  `json:"code" binding:"required"`
	ProductName string  `json:"product_name" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

type CreateOrderRequest struct {
	UserID          string            `json:"user_id" binding:"required"`
	Products        []*OrderedProduct `json:"products" binding:"required,min=1"`
	ShippingAddress string            `json:"shipping_address" binding:"required"`
	PaymentMethod   string            `json:"payment_method" binding:"required"`
}

type OrderSummaryResponse struct {
	OrderID         string            `json:"order_id"`
	UserID          string            `json:"user_id"`
	Status          enums.OrderStatus `json:"status"`
	TotalAmount     float64           `json:"total_amount"`
	ShippingAddress string            `json:"shipping_address"`
	PaymentMethod   string            `json:"payment_method"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
}

type OrderResponse struct {
	OrderID         string            `json:"order_id"`
	UserID          string            `json:"user_id"`
	Status          enums.OrderStatus `json:"status"`
	TotalAmount     float64           `json:"total_amount"`
	ShippingAddress string            `json:"shipping_address"`
	PaymentMethod   string            `json:"payment_method"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	Products        []*OrderedProduct `json:"products"`
}

type UserOrdersResponse struct {
	UserID string                  `json:"user_id"`
	Orders []*OrderSummaryResponse `json:"orders"`
}

type UpdateOrderStatusRequest struct {
	Status enums.OrderStatus `json:"status" binding:"required"`
}
