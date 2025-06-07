package repository

import (
	"errors"
	"sync"

	"github.com/dinosgnk/agora-project/internal/services/cart/model"
)

type ICartRepository interface {
	GetCart(userId string) (*model.Cart, error)
	UpdateCart(cart *model.Cart) error
	Clear(userId string) error
}

type CartMap struct {
	data map[string]*model.Cart
	mu   sync.RWMutex
}

func NewCartRepository() *CartMap {
	return &CartMap{
		data: make(map[string]*model.Cart),
	}
}

func (cm *CartMap) GetCart(userId string) (*model.Cart, error) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	if cart, ok := cm.data[userId]; ok {
		return cart, nil
	}
	return nil, errors.New("Cart not found")
}

func (cm *CartMap) UpdateCart(cart *model.Cart) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.data[cart.UserId] = cart
	return nil
}

func (cm *CartMap) Clear(userId string) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	delete(cm.data, userId)
	return nil
}
