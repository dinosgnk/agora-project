package repository

import (
	"fmt"

	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	"github.com/dinosgnk/agora-project/internal/pkg/postgres"
	"github.com/dinosgnk/agora-project/internal/services/order/enums"
	"github.com/dinosgnk/agora-project/internal/services/order/model"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type PostgresOrderRepository struct {
	gormDb *postgres.GormDatabase
}

func NewPostgresOrderRepository(logger logger.Logger) *PostgresOrderRepository {
	gormDb, err := postgres.NewGormDatabase(
		logger,
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "orders.t_",
				SingularTable: true,
			},
		},
	)

	if err != nil {
		return nil
	}

	return &PostgresOrderRepository{
		gormDb: gormDb,
	}
}

func (repo *PostgresOrderRepository) CreateOrder(order *model.Order, products []*model.OrderedProduct) (*model.Order, error) {
	tx := repo.gormDb.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, product := range products {
		if err := tx.Create(product).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return repo.GetOrderSummaryByID(order.ID)
}

func (repo *PostgresOrderRepository) GetAllOrderSummaries() ([]*model.Order, error) {
	var orders []*model.Order
	result := repo.gormDb.Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (repo *PostgresOrderRepository) GetAllOrders() ([]*model.OrderWithProducts, error) {
	var orders []*model.Order
	result := repo.gormDb.Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	// For each order, fetch its products
	ordersWithProducts := make([]*model.OrderWithProducts, 0, len(orders))
	for _, order := range orders {
		products, err := repo.GetProductsByOrderID(order.ID)
		if err != nil {
			return nil, err
		}
		ordersWithProducts = append(ordersWithProducts, &model.OrderWithProducts{
			Order:    *order,
			Products: products,
		})
	}

	return ordersWithProducts, nil
}

func (repo *PostgresOrderRepository) GetOrderSummaryByID(orderId string) (*model.Order, error) {
	var order model.Order
	result := repo.gormDb.Where("id = ?", orderId).First(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (repo *PostgresOrderRepository) GetOrderByID(orderId string) (*model.OrderWithProducts, error) {
	order, err := repo.GetOrderSummaryByID(orderId)
	if err != nil {
		return nil, err
	}

	products, err := repo.GetProductsByOrderID(orderId)
	if err != nil {
		return nil, err
	}

	return &model.OrderWithProducts{
		Order:    *order,
		Products: products,
	}, nil
}

func (repo *PostgresOrderRepository) GetAllOrderSummariesByUserID(userId string) ([]*model.Order, error) {
	var orders []*model.Order
	result := repo.gormDb.Where("user_id = ?", userId).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (repo *PostgresOrderRepository) GetAllOrdersByUserID(userId string) ([]*model.OrderWithProducts, error) {
	orders, err := repo.GetAllOrderSummariesByUserID(userId)
	if err != nil {
		return nil, err
	}

	ordersWithProducts := make([]*model.OrderWithProducts, 0, len(orders))
	for _, order := range orders {
		products, err := repo.GetProductsByOrderID(order.ID)
		if err != nil {
			return nil, err
		}
		ordersWithProducts = append(ordersWithProducts, &model.OrderWithProducts{
			Order:    *order,
			Products: products,
		})
	}

	return ordersWithProducts, nil
}

func (repo *PostgresOrderRepository) GetProductsByOrderID(orderId string) ([]*model.OrderedProduct, error) {
	var orderedProducts []*model.OrderedProduct
	result := repo.gormDb.Where("order_id = ?", orderId).Find(&orderedProducts)
	if result.Error != nil {
		return nil, result.Error
	}

	return orderedProducts, nil
}

func (repo *PostgresOrderRepository) UpdateOrderStatus(orderId string, status enums.OrderStatus) error {
	result := repo.gormDb.Model(&model.Order{}).Where("id = ?", orderId).Update("status", status)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("order with id %s not found", orderId)
	}

	return nil
}
