package repository

import (
	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	"github.com/dinosgnk/agora-project/internal/pkg/postgres"
	"github.com/dinosgnk/agora-project/internal/services/catalog/model"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type PostgresProductRepository struct {
	gormDb *postgres.GormDatabase
}

func NewPostgresProductRepository(logger logger.Logger) *PostgresProductRepository {
	gormDb, err := postgres.NewGormDatabase(
		logger,
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "products.t_",
				SingularTable: true,
			},
		},
	)

	if err != nil {
		return nil
	}

	return &PostgresProductRepository{
		gormDb: gormDb,
	}
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

func (repo *PostgresProductRepository) GetProductByCode(productCode string) (*model.Product, error) {
	var product *model.Product
	result := repo.gormDb.Where("product_code = ?", productCode).First(&product)
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
	err := repo.gormDb.Model(&model.Product{}).Where("product_code = ?", product.ProductCode).Updates(product).Error
	if err != nil {
		return product, err
	}

	return product, err
}

func (repo *PostgresProductRepository) DeleteProduct(productCode string) (bool, error) {
	result := repo.gormDb.Delete(&model.Product{}, "product_code = ?", productCode)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
