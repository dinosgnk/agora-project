package repository

import (
	"errors"
	"sync"

	"github.com/dinosgnk/agora-project/internal/services/cart/model"
)

type MockCartRepository struct {
	data map[string]*model.Cart
	mu   sync.RWMutex
}

func NewMockCartRepository() *MockCartRepository {
	return &MockCartRepository{
		data: make(map[string]*model.Cart),
	}
}

func (cm *MockCartRepository) GetCartByUserId(userId string) (*model.Cart, error) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	if cart, ok := cm.data[userId]; ok {
		return cart, nil
	}
	return nil, errors.New("Cart not found")
}

func (cm *MockCartRepository) UpdateCart(cart *model.Cart) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.data[cart.UserId] = cart
	return nil
}

func (cm *MockCartRepository) Clear(userId string) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	delete(cm.data, userId)
	return nil
}
