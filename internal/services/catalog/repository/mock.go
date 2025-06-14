package repository

import (
	"errors"

	"github.com/dinosgnk/agora-project/internal/services/catalog/model"
)

type MockProductRepository struct {
	data map[string]*model.Product
}

func NewMockProductRepository() *MockProductRepository {
	return &MockProductRepository{
		data: make(map[string]*model.Product),
	}
}

func (repo *MockProductRepository) GetAllProducts() ([]*model.Product, error) {
	var productList []*model.Product
	for _, product := range repo.data {
		productList = append(productList, product)
	}
	return productList, nil
}

func (repo *MockProductRepository) GetProductsByCategory(category string) ([]*model.Product, error) {
	var productList []*model.Product
	for _, product := range repo.data {
		if product.Category == category {
			productList = append(productList, product)
		}
	}
	if len(productList) == 0 {
		return nil, errors.New("no products found in this category")
	}
	return productList, nil
}

func (repo *MockProductRepository) CreateProduct(product *model.Product) error {
	if _, exists := repo.data[product.ProductId]; exists {
		return errors.New("product already exists")
	}
	repo.data[product.ProductId] = product
	return nil
}

func (repo *MockProductRepository) GetProduct(id string) (*model.Product, error) {
	product, exists := repo.data[id]
	if !exists {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (repo *MockProductRepository) UpdateProduct(product *model.Product) error {
	if _, exists := repo.data[product.ProductId]; !exists {
		return errors.New("product not found")
	}
	repo.data[product.ProductId] = product
	return nil
}

func (repo *MockProductRepository) DeleteProduct(id string) error {
	if _, exists := repo.data[id]; !exists {
		return errors.New("product not found")
	}
	delete(repo.data, id)
	return nil
}
