package repository

import (
	"fmt"

	"github.com/dinosgnk/agora-project/internal/services/catalog/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type PostgresProductRepository struct {
	db *gorm.DB
}

func NewPostgresProductRepository() *PostgresProductRepository {
	datasource := "postgres://devuser:devpass@localhost:5432/AgoraDB?sslmode=disable"

	gormDb, err := gorm.Open(postgres.Open(datasource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "products.t_",
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Printf("Cannot connect to database")
	}

	// Set up connection pool
	sqlDB, err := gormDb.DB()
	if err != nil {

	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(30)

	return &PostgresProductRepository{db: gormDb}
}

func (repo *PostgresProductRepository) GetAllProducts() ([]*model.Product, error) {
	var products []*model.Product
	result := repo.db.Find(products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (repo *PostgresProductRepository) GetProductsByCategory(category string) ([]*model.Product, error) {
	var products []*model.Product
	result := repo.db.Where("category = ?", category).Find(products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (repo *PostgresProductRepository) GetProductById(id string) (*model.Product, error) {
	var product *model.Product
	result := repo.db.Where("product_id = ?", id).Find(product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repo *PostgresProductRepository) CreateProduct(product *model.Product) (*model.Product, error) {
	result := repo.db.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *PostgresProductRepository) UpdateProduct(product *model.Product) (*model.Product, error) {
	err := repo.db.Model(&model.Product{}).Where("product_id = ?", product.ProductId).Updates(product).Error
	return product, err
}

func (repo *PostgresProductRepository) DeleteProduct(id string) (bool, error) {
	result := repo.db.Delete(&model.Product{}, "prodict_id = ?", id)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
