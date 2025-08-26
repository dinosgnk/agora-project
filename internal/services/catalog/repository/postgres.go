package repository

import (
	"fmt"

	"github.com/dinosgnk/agora-project/internal/services/catalog/model"
	"github.com/dinosgnk/agora-project/pkg/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type PostgresProductRepository struct {
	gormDb *postgres.GormDatabase
}

func NewPostgresProductRepository() *PostgresProductRepository {

	gormDb, err := postgres.NewGormDatabase(&gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "products.t_",
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Printf("Cannot connect to database")
	}

	return &PostgresProductRepository{gormDb: gormDb}
}

func (repo *PostgresProductRepository) GetAllProducts() ([]*model.Product, error) {
	var products []*model.Product
	result := repo.gormDb.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (repo *PostgresProductRepository) GetProductsByCategory(category string) ([]*model.Product, error) {
	var products []*model.Product
	result := repo.gormDb.Where("category = ?", category).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (repo *PostgresProductRepository) GetProductById(id string) (*model.Product, error) {
	var product *model.Product
	result := repo.gormDb.Where("product_id = ?", id).Find(product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repo *PostgresProductRepository) CreateProduct(product *model.Product) (*model.Product, error) {
	result := repo.gormDb.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *PostgresProductRepository) UpdateProduct(product *model.Product) (*model.Product, error) {
	err := repo.gormDb.Model(&model.Product{}).Where("product_id = ?", product.ProductId).Updates(product).Error
	return product, err
}

func (repo *PostgresProductRepository) DeleteProduct(id string) (bool, error) {
	result := repo.gormDb.Delete(&model.Product{}, "prodict_id = ?", id)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
