package repository

import (
	"github.com/dinosgnk/agora-project/internal/services/cart/model"
)

type ICartRepository interface {
	GetCartByUserId(userId string) (*model.Cart, error)
	UpdateCart(cart *model.Cart) error
	Clear(userId string) error
}
