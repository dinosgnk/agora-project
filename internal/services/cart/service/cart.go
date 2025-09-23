package service

import (
	"errors"

	"github.com/dinosgnk/agora-project/internal/services/cart/dto"
	"github.com/dinosgnk/agora-project/internal/services/cart/model"
	"github.com/dinosgnk/agora-project/internal/services/cart/repository"
)

type ICartService interface {
	GetCartByUserId(userId string) (*dto.CartResponse, error)
	AddItem(userId string, itemToAdd *dto.Item) error
	RemoveItem(userId string, productCode string) error
	UpdateCart(userId string, updatedCart map[string]int) error
	ClearCart(userId string) error
}

type CartService struct {
	repo repository.ICartRepository
}

func NewCartService(repo repository.ICartRepository) *CartService {
	return &CartService{
		repo: repo,
	}
}

func (cs *CartService) GetCartByUserId(userId string) (*dto.CartResponse, error) {
	cart, err := cs.repo.GetCartByUserId(userId)
	if err != nil {
		return nil, err
	}

	return cs.mapCartModelToDto(cart), nil
}

func (cs *CartService) AddItem(userId string, itemToAdd *dto.Item) error {
	cart, err := cs.repo.GetCartByUserId(userId)
	if err != nil {
		// Cart doesn't exist, create a new one
		cart = &model.Cart{
			UserId: userId,
			Items:  []*model.Item{},
		}
	}

	for i, item := range cart.Items {
		if item.ProductCode == itemToAdd.ProductCode {
			cart.Items[i].Quantity += 1
			return cs.repo.UpdateCart(cart)
		}
	}

	newItem := cs.mapItemDtoToModel(itemToAdd)
	cart.Items = append(cart.Items, newItem)
	return cs.repo.UpdateCart(cart)
}

func (cs *CartService) RemoveItem(userId string, productCode string) error {
	cart, err := cs.repo.GetCartByUserId(userId)
	if err != nil {
		return errors.New("cart not found")
	}

	// Filter out the item
	newItems := make([]*model.Item, 0, len(cart.Items))
	for _, item := range cart.Items {
		if item.ProductCode != productCode {
			newItems = append(newItems, item)
		}
	}

	cart.Items = newItems
	return cs.repo.UpdateCart(cart)
}

func (cs *CartService) UpdateCart(userId string, updatedCart map[string]int) error {
	cart, err := cs.repo.GetCartByUserId(userId)
	if err != nil {
		return errors.New("cart not found")
	}

	for i, item := range cart.Items {
		if newQuantity, exists := updatedCart[item.ProductCode]; exists {
			if newQuantity == 0 {
				cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
				i--
			}
			cart.Items[i].Quantity = newQuantity
		}
	}

	return cs.repo.UpdateCart(cart)
}

func (cs *CartService) ClearCart(userId string) error {
	err := cs.repo.Clear(userId)
	if err != nil {
		return err
	}

	return nil
}

// Helper functions to map between DTOs and Models
func (cs *CartService) mapCartModelToDto(cart *model.Cart) *dto.CartResponse {
	items := make([]dto.Item, len(cart.Items))
	for i, item := range cart.Items {
		items[i] = cs.mapItemModelToDto(item)
	}

	return &dto.CartResponse{
		UserId: cart.UserId,
		Items:  items,
	}
}

func (cs *CartService) mapItemModelToDto(item *model.Item) dto.Item {
	return dto.Item{
		ProductCode: item.ProductCode,
		Name:        item.Name,
		Quantity:    item.Quantity,
		Price:       item.Price,
	}
}

func (cs *CartService) mapItemDtoToModel(item *dto.Item) *model.Item {
	return &model.Item{
		ProductCode: item.ProductCode,
		Name:        item.Name,
		Quantity:    item.Quantity,
		Price:       item.Price,
	}
}
