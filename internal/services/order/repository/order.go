package repository

import (
	"github.com/dinosgnk/agora-project/internal/services/order/enums"
	"github.com/dinosgnk/agora-project/internal/services/order/model"
)

type IOrderRepository interface {
	CreateOrder(order *model.Order, products []*model.OrderedProduct) (*model.Order, error)
	GetAllOrderSummaries() ([]*model.Order, error)
	GetAllOrders() ([]*model.OrderWithProducts, error)
	GetOrderSummaryByID(orderId string) (*model.Order, error)
	GetOrderByID(orderId string) (*model.OrderWithProducts, error)
	GetAllOrderSummariesByUserID(userId string) ([]*model.Order, error)
	GetAllOrdersByUserID(userId string) ([]*model.OrderWithProducts, error)
	GetProductsByOrderID(orderId string) ([]*model.OrderedProduct, error)
	UpdateOrderStatus(orderId string, status enums.OrderStatus) error
}
