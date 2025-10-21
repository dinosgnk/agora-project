package messaging

import "time"

type OrderEvent struct {
	EventID   string    `json:"event_id"`
	Timestamp time.Time `json:"timestamp"`
	OrderID   string    `json:"order_id"`
	UserID    string    `json:"user_id"`
}

type OrderCreatedEvent struct {
	OrderEvent
	TotalAmount     float64               `json:"total_amount"`
	ShippingAddress string                `json:"shipping_address"`
	PaymentMethod   string                `json:"payment_method"`
	Products        []OrderCreatedProduct `json:"products"`
}

type OrderCreatedProduct struct {
	ProductCode string  `json:"product_code"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

type OrderStatusUpdatedEvent struct {
	OrderEvent
	OldStatus string `json:"old_status"`
	NewStatus string `json:"new_status"`
}

type OrderConfirmedEvent struct {
	OrderEvent
	PaymentMethod string  `json:"payment_method"`
	TotalAmount   float64 `json:"total_amount"`
}

type OrderProcessingEvent struct {
	OrderEvent
}

type OrderShippedEvent struct {
	OrderEvent
	TrackingNumber string `json:"tracking_number"`
}

type OrderDeliveredEvent struct {
	OrderEvent
}

type OrderCancelledEvent struct {
	OrderEvent
	Reason string `json:"reason"`
}
