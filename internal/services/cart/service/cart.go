package service

import (
	"errors"

	"github.com/dinosgnk/agora-project/internal/services/cart/model"
	"github.com/dinosgnk/agora-project/internal/services/cart/repository"
)

type ICartService interface {
	GetCart(userId string) (*model.Cart, error)
	AddItem(userId string, itemToAdd *model.Item) error
	RemoveItem(userId string, productId string) error
	UpdateCart(userId string, updatedCart map[string]int) error
	ClearCart(userId string) error
}

type CartService struct {
	repo repository.ICartRepository
}

func NewCartService(repo repository.ICartRepository) *CartService {
	return &CartService{repo: repo}
}

func (cs *CartService) GetCart(userId string) (*model.Cart, error) {
	return cs.repo.GetCart(userId)
}

func (cs *CartService) AddItem(userId string, itemToAdd *model.Item) error {
	cart, err := cs.repo.GetCart(userId)
	if err != nil {
		cart = &model.Cart{
			UserId: userId,
			Items:  []*model.Item{},
		}
	}

	// Check if item already exists
	for i, item := range cart.Items {
		if item.ProductId == itemToAdd.ProductId {
			cart.Items[i].Quantity += 1
			return cs.repo.UpdateCart(cart)
		}
	}

	cart.Items = append(cart.Items, itemToAdd)
	return cs.repo.UpdateCart(cart)
}

func (cs *CartService) RemoveItem(userId string, productId string) error {
	cart, err := cs.repo.GetCart(userId)
	if err != nil {
		return errors.New("Cart not found")
	}

	// Filter out the item
	newItems := make([]*model.Item, 0, len(cart.Items))
	for _, item := range cart.Items {
		if item.ProductId != productId {
			newItems = append(newItems, item)
		}
	}

	cart.Items = newItems
	return cs.repo.UpdateCart(cart)
}

func (cs *CartService) UpdateCart(userId string, updatedCart map[string]int) error {
	cart, err := cs.repo.GetCart(userId)
	if err != nil {
		return errors.New("Cart not found")
	}

	for i, item := range cart.Items {
		if newQuantity, exists := updatedCart[item.ProductId]; exists {
			if newQuantity < 0 {
				return errors.New("Invalid quantity for product: " + item.ProductId)
			} else if newQuantity == 0 {
				cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
				i--
			}
			cart.Items[i].Quantity = newQuantity
		}
	}

	return cs.repo.UpdateCart(cart)
}

func (cs *CartService) ClearCart(userId string) error {
	return cs.repo.Clear(userId)
}
