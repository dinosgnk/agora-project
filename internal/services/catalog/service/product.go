package service

import (
	"github.com/dinosgnk/agora-project/internal/services/catalog/model"
	"github.com/dinosgnk/agora-project/internal/services/catalog/repository"
)

type IProductService interface {
	GetAllProducts() (*[]model.Product, error)
	GetProductsByCategory(category string) (*[]model.Product, error)
	GetProductById(id string) (*model.Product, error)
	CreateProduct(model.Product) (*model.Product, error)
	UpdateProduct(id string)
	DeleteProduct(id string)
}

type ProductService struct {
	repo repository.IProductRepository
}

func NewProductService(repo repository.IProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (p *ProductService) GetAllProducts() (*[]model.Product, error) {
	products, err := p.repo.GetAllProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductService) GetProductsByCategory(category string) (*[]model.Product, error) {
	products, err := p.repo.GetProductsByCategory(category)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductService) GetProductById(id string) (*model.Product, error) {
	product, err := p.repo.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductService) CreateProduct(product *model.Product) (*model.Product, error) {
	createdProduct, err := p.repo.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return createdProduct, nil
}

func (p *ProductService) UpdateProduct(id string, updatedProduct *model.Product) (*model.Product, error) {
	product, err := p.repo.UpdateProduct(id, updatedProduct)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) DeleteProduct(id string) (bool, error) {
	productDeleted, err := s.repo.DeleteProduct(id)
	if err != nil {
		return false, err
	}

	return productDeleted, nil
}
