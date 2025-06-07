package service

import (
	"agora/basket/model"
	"agora/basket/repository"
)

type IBasketService interface {
	GetBasket(userId string) (*model.Basket, error)
	AddItem(userId string, item model.Item) error
	ClearBasket(userId string) error
}

type BasketService struct {
	repo repository.IBasketRepository
}

func NewBasketService(repo repository.BasketRepository) *BasketService {
	return &basketService{repo: repo}
}

func (b *BasketService) GetBasket(userId string) (*model.Basket, error) {
	return s.repo.Get(userId)
}

func (b *BasketService) AddItem(userId string, item model.Item) error {
	basket, err := s.repo.Get(userId)
	if err != nil {
		basket = &model.Basket{
			userId: userId,
			Items:  []model.Item{},
		}
	}

	// Check if item already exists
	for i, it := range basket.Items {
		if it.ProductID == item.ProductID {
			basket.Items[i].Quantity += item.Quantity
			return s.repo.Update(basket)
		}
	}

	basket.Items = append(basket.Items, item)
	return s.repo.Update(basket)
}

func (b *BasketService) ClearBasket(userId string) error {
	return s.repo.Clear(userId)
}
