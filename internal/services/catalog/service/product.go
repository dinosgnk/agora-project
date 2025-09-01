package service

import (
	"github.com/dinosgnk/agora-project/internal/services/catalog/dto"
	"github.com/dinosgnk/agora-project/internal/services/catalog/model"
	"github.com/dinosgnk/agora-project/internal/services/catalog/repository"
)

type IProductService interface {
	GetAllProducts() ([]*model.Product, error)
	GetProductsByCategory(category string) ([]*model.Product, error)
	GetProductByCode(productCode string) (*dto.ProductResponse, error)
	CreateProduct(productReq *dto.CreateProductRequest) (*dto.ProductResponse, error)
	UpdateProduct(productCode string, product *model.Product) (*model.Product, error)
	DeleteProduct(productCode string) (bool, error)
}

type ProductService struct {
	repo repository.IProductRepository
}

func NewProductService(repo repository.IProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (p *ProductService) GetAllProducts() ([]*model.Product, error) {
	products, err := p.repo.GetAllProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductService) GetProductsByCategory(category string) ([]*model.Product, error) {
	products, err := p.repo.GetProductsByCategory(category)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductService) GetProductByCode(productCode string) (*dto.ProductResponse, error) {
	product, err := p.repo.GetProductByCode(productCode)
	if err != nil {
		return nil, err
	}

	return p.mapProductModelToDto(product), nil
}

func (p *ProductService) CreateProduct(productReq *dto.CreateProductRequest) (*dto.ProductResponse, error) {
	product := p.mapProductDtoToModel(productReq)

	createdProduct, err := p.repo.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return p.mapProductModelToDto(createdProduct), nil
}

func (p *ProductService) UpdateProduct(productCode string, updatedProduct *model.Product) (*model.Product, error) {
	product, err := p.repo.UpdateProduct(updatedProduct)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) DeleteProduct(productCode string) (bool, error) {
	productDeleted, err := s.repo.DeleteProduct(productCode)
	if err != nil {
		return false, err
	}

	return productDeleted, nil
}

// Helper functions to map between DTOs and Models
func (p *ProductService) mapProductDtoToModel(dto *dto.CreateProductRequest) *model.Product {
	return &model.Product{
		ProductCode: dto.ProductCode,
		Name:        dto.Name,
		Category:    dto.Category,
		Description: dto.Description,
		Price:       dto.Price,
	}
}

func (p *ProductService) mapProductModelToDto(product *model.Product) *dto.ProductResponse {
	return &dto.ProductResponse{
		ProductCode: product.ProductCode,
		Name:        product.Name,
		Category:    product.Category,
		Description: product.Description,
		Price:       product.Price,
	}
}
