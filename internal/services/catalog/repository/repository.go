package repository

import (
	"github.com/dinosgnk/agora-project/internal/services/catalog/model"
)

type IProductRepository interface {
	GetAllProducts() (*[]model.Product, error)
	GetProductsByCategory(category string) (*[]model.Product, error)
	GetProductById(id string) (*model.Product, error)
	CreateProduct(*model.Product) (*model.Product, error)
	UpdateProduct(*model.Product) (*model.Product, error)
	DeleteProduct(id string) (bool, error)
}
