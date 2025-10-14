package repository

import (
	"fmt"
	"sync"

	"github.com/dinosgnk/agora-project/internal/services/order/enums"
	"github.com/dinosgnk/agora-project/internal/services/order/model"
)

type MockOrderRepository struct {
	orders   map[string]*model.Order
	products map[string][]*model.OrderedProduct
	mutex    sync.RWMutex
}

func NewMockOrderRepository() *MockOrderRepository {
	return &MockOrderRepository{
		orders:   make(map[string]*model.Order),
		products: make(map[string][]*model.OrderedProduct),
	}
}

func (repo *MockOrderRepository) CreateOrder(order *model.Order, products []*model.OrderedProduct) (*model.Order, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	repo.orders[order.ID] = order
	repo.products[order.ID] = products
	return order, nil
}

func (repo *MockOrderRepository) GetAllOrderSummaries() ([]*model.Order, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	var allOrders []*model.Order
	for _, order := range repo.orders {
		allOrders = append(allOrders, order)
	}

	return allOrders, nil
}

func (repo *MockOrderRepository) GetAllOrders() ([]*model.OrderWithProducts, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	var allOrders []*model.OrderWithProducts
	for _, order := range repo.orders {
		products := repo.products[order.ID]

		allOrders = append(allOrders, &model.OrderWithProducts{
			Order:    *order,
			Products: products,
		})
	}

	return allOrders, nil
}

func (repo *MockOrderRepository) GetOrderSummaryByID(orderId string) (*model.Order, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	order, exists := repo.orders[orderId]
	if !exists {
		return nil, fmt.Errorf("order with id %s not found", orderId)
	}

	return order, nil
}

func (repo *MockOrderRepository) GetOrderByID(orderId string) (*model.OrderWithProducts, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	order, exists := repo.orders[orderId]
	if !exists {
		return nil, fmt.Errorf("order with id %s not found", orderId)
	}

	products := repo.products[orderId]

	return &model.OrderWithProducts{
		Order:    *order,
		Products: products,
	}, nil
}

func (repo *MockOrderRepository) GetAllOrderSummariesByUserID(userId string) ([]*model.Order, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	var userOrders []*model.Order
	for _, order := range repo.orders {
		if order.UserID == userId {
			userOrders = append(userOrders, order)
		}
	}

	return userOrders, nil
}

func (repo *MockOrderRepository) GetAllOrdersByUserID(userId string) ([]*model.OrderWithProducts, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	var userOrders []*model.OrderWithProducts
	for _, order := range repo.orders {
		if order.UserID == userId {
			products := repo.products[order.ID]

			userOrders = append(userOrders, &model.OrderWithProducts{
				Order:    *order,
				Products: products,
			})
		}
	}

	return userOrders, nil
}

func (repo *MockOrderRepository) GetProductsByOrderID(orderId string) ([]*model.OrderedProduct, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	products, exists := repo.products[orderId]
	if !exists {
		return []*model.OrderedProduct{}, nil
	}

	return products, nil
}

func (repo *MockOrderRepository) UpdateOrderStatus(orderId string, status enums.OrderStatus) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	order, exists := repo.orders[orderId]
	if !exists {
		return fmt.Errorf("order with id %s not found", orderId)
	}

	order.Status = status
	return nil
}
