package repository

import (
	"errors"
	"sync"

	"agora/basket/model"
)

type BasketRepository interface {
	Get(userId string) (*model.Basket, error)
	Update(basket *model.Basket) error
	Clear(userId string) error
}

type BasketMap struct {
	data map[string]*model.Basket
	mu   sync.RWMutex
}

func NewBasketMap() BasketRepository {
	return &BasketMap{
		data: make(map[string]*model.Basket),
	}
}

func (bm *BasketMap) Get(userId string) (*model.Basket, error) {
	bm.mu.RLock()
	defer bm.mu.RUnlock()

	if basket, ok := bm.data[userId]; ok {
		return basket, nil
	}
	return nil, errors.New("Basket not found")
}

func (bm *BasketMap) Update(basket *model.Basket) error {
	bm.mu.Lock()
	defer bm.mu.Unlock()

	r.data[basket.UserId] = basket
	return nil
}

func (bm *BasketMap) Clear(userId string) error {
	bm.mu.Lock()
	defer bm.mu.Unlock()

	delete(dm.data, userId)
	return nil
}
