package service

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/dinosgnk/agora-project/internal/services/order/dto"
	"github.com/dinosgnk/agora-project/internal/services/order/enums"
	"github.com/dinosgnk/agora-project/internal/services/order/model"
	"github.com/dinosgnk/agora-project/internal/services/order/repository"
)

type IOrderService interface {
	CreateOrder(orderReq *dto.CreateOrderRequest) (*dto.OrderResponse, error)
	GetAllOrderSummaries() ([]*dto.OrderSummaryResponse, error)
	GetAllOrders() ([]*dto.OrderResponse, error)
	GetOrderSummaryByID(orderId string) (*dto.OrderSummaryResponse, error)
	GetOrderByID(orderId string) (*dto.OrderResponse, error)
	GetAllOrderSummariesByUserID(userId string) ([]*dto.OrderSummaryResponse, error)
	GetAllOrdersByUserID(userId string) ([]*dto.OrderResponse, error)
	GetProductsByOrderID(orderId string) ([]*dto.OrderedProduct, error)
	UpdateOrderStatus(orderId string, statusReq *dto.UpdateOrderStatusRequest) error
}

type OrderService struct {
	repo repository.IOrderRepository
}

func NewOrderService(repo repository.IOrderRepository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) CreateOrder(orderReq *dto.CreateOrderRequest) (*dto.OrderResponse, error) {
	orderId := uuid.New().String()

	var totalAmount float64
	var orderProducts = make([]*model.OrderedProduct, 0, len(orderReq.Products))
	for _, product := range orderReq.Products {
		subtotal := product.Price * float64(product.Quantity)
		totalAmount += subtotal

		orderProducts = append(orderProducts, &model.OrderedProduct{
			ID:          uuid.New().String(),
			OrderID:     orderId,
			ProductCode: product.ProductCode,
			ProductName: product.ProductName,
			Quantity:    product.Quantity,
			Price:       product.Price,
			Subtotal:    subtotal,
		})
	}

	order := &model.Order{
		ID:              orderId,
		UserID:          orderReq.UserID,
		Status:          enums.OrderStatusPending,
		TotalAmount:     totalAmount,
		ShippingAddress: orderReq.ShippingAddress,
		PaymentMethod:   orderReq.PaymentMethod,
	}

	createdOrder, err := s.repo.CreateOrder(order, orderProducts)
	if err != nil {
		return nil, err
	}

	return &dto.OrderResponse{
		OrderID:         createdOrder.ID,
		UserID:          createdOrder.UserID,
		Status:          createdOrder.Status,
		TotalAmount:     createdOrder.TotalAmount,
		ShippingAddress: createdOrder.ShippingAddress,
		PaymentMethod:   createdOrder.PaymentMethod,
		CreatedAt:       createdOrder.CreatedAt,
		UpdatedAt:       createdOrder.UpdatedAt,
		Products:        orderReq.Products,
	}, nil

}

func (s *OrderService) GetAllOrderSummaries() ([]*dto.OrderSummaryResponse, error) {
	orders, err := s.repo.GetAllOrderSummaries()
	if err != nil {
		return nil, err
	}

	var ordersSummary []*dto.OrderSummaryResponse
	for _, order := range orders {
		ordersSummary = append(ordersSummary, &dto.OrderSummaryResponse{
			OrderID:         order.ID,
			UserID:          order.UserID,
			Status:          order.Status,
			TotalAmount:     order.TotalAmount,
			ShippingAddress: order.ShippingAddress,
			PaymentMethod:   order.PaymentMethod,
			CreatedAt:       order.CreatedAt,
			UpdatedAt:       order.UpdatedAt,
		})
	}

	return ordersSummary, nil
}

func (s *OrderService) GetAllOrders() ([]*dto.OrderResponse, error) {
	ordersWithProducts, err := s.repo.GetAllOrders()
	if err != nil {
		return nil, err
	}

	var orderResponse []*dto.OrderResponse
	for _, order := range ordersWithProducts {
		orderedProducts := make([]*dto.OrderedProduct, 0, len(order.Products))
		for _, p := range order.Products {
			orderedProducts = append(orderedProducts, &dto.OrderedProduct{
				ProductCode: p.ProductCode,
				ProductName: p.ProductName,
				Quantity:    p.Quantity,
				Price:       p.Price,
			})
		}

		orderResponse = append(orderResponse, &dto.OrderResponse{
			OrderID:         order.Order.ID,
			UserID:          order.Order.UserID,
			Status:          order.Order.Status,
			TotalAmount:     order.Order.TotalAmount,
			ShippingAddress: order.Order.ShippingAddress,
			PaymentMethod:   order.Order.PaymentMethod,
			CreatedAt:       order.Order.CreatedAt,
			UpdatedAt:       order.Order.UpdatedAt,
			Products:        orderedProducts,
		})
	}

	return orderResponse, nil
}

func (s *OrderService) GetOrderSummaryByID(orderId string) (*dto.OrderSummaryResponse, error) {
	order, err := s.repo.GetOrderSummaryByID(orderId)
	if err != nil {
		return nil, err
	}

	return &dto.OrderSummaryResponse{
		OrderID:         order.ID,
		UserID:          order.UserID,
		Status:          order.Status,
		TotalAmount:     order.TotalAmount,
		ShippingAddress: order.ShippingAddress,
		PaymentMethod:   order.PaymentMethod,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       order.UpdatedAt,
	}, nil

}

func (s *OrderService) GetOrderByID(orderId string) (*dto.OrderResponse, error) {
	order, err := s.repo.GetOrderByID(orderId)
	if err != nil {
		return nil, err
	}

	products, err := s.repo.GetProductsByOrderID(orderId)
	if err != nil {
		return nil, err
	}

	var orderedProducts []*dto.OrderedProduct
	for _, p := range products {
		orderedProducts = append(orderedProducts, &dto.OrderedProduct{
			ProductCode: p.ProductCode,
			ProductName: p.ProductName,
			Quantity:    p.Quantity,
			Price:       p.Price,
		})
	}

	return &dto.OrderResponse{
		OrderID:         order.Order.ID,
		UserID:          order.Order.UserID,
		Status:          order.Order.Status,
		TotalAmount:     order.Order.TotalAmount,
		ShippingAddress: order.Order.ShippingAddress,
		PaymentMethod:   order.Order.PaymentMethod,
		CreatedAt:       order.Order.CreatedAt,
		UpdatedAt:       order.Order.UpdatedAt,
		Products:        orderedProducts,
	}, nil

}

func (s *OrderService) GetAllOrderSummariesByUserID(userId string) ([]*dto.OrderSummaryResponse, error) {
	orders, err := s.repo.GetAllOrderSummariesByUserID(userId)
	if err != nil {
		return nil, err
	}

	var orderSummaries []*dto.OrderSummaryResponse
	for _, order := range orders {
		orderSummaries = append(orderSummaries, &dto.OrderSummaryResponse{
			OrderID:         order.ID,
			UserID:          order.UserID,
			Status:          order.Status,
			TotalAmount:     order.TotalAmount,
			ShippingAddress: order.ShippingAddress,
			PaymentMethod:   order.PaymentMethod,
			CreatedAt:       order.CreatedAt,
			UpdatedAt:       order.UpdatedAt,
		})
	}

	return orderSummaries, nil
}

func (s *OrderService) GetAllOrdersByUserID(userId string) ([]*dto.OrderResponse, error) {
	orders, err := s.repo.GetAllOrdersByUserID(userId)
	if err != nil {
		return nil, err
	}

	orderResponse := make([]*dto.OrderResponse, 0, len(orders))
	for _, order := range orders {

		orderedProducts := make([]*dto.OrderedProduct, 0, len(order.Products))
		for _, p := range order.Products {
			orderedProducts = append(orderedProducts, &dto.OrderedProduct{
				ProductCode: p.ProductCode,
				ProductName: p.ProductName,
				Quantity:    p.Quantity,
				Price:       p.Price,
			})
		}

		orderResponse = append(orderResponse, &dto.OrderResponse{
			OrderID:         order.Order.ID,
			UserID:          order.Order.UserID,
			Status:          order.Order.Status,
			TotalAmount:     order.Order.TotalAmount,
			ShippingAddress: order.Order.ShippingAddress,
			PaymentMethod:   order.Order.PaymentMethod,
			CreatedAt:       order.Order.CreatedAt,
			UpdatedAt:       order.Order.UpdatedAt,
			Products:        orderedProducts,
		})
	}

	return orderResponse, nil
}

func (s *OrderService) GetProductsByOrderID(orderId string) ([]*dto.OrderedProduct, error) {
	products, err := s.repo.GetProductsByOrderID(orderId)
	if err != nil {
		return nil, err
	}

	var orderProducts []*dto.OrderedProduct
	for _, p := range products {
		orderProducts = append(orderProducts, &dto.OrderedProduct{
			ProductCode: p.ProductCode,
			ProductName: p.ProductName,
			Quantity:    p.Quantity,
			Price:       p.Price,
		})
	}

	return orderProducts, nil
}

func (s *OrderService) UpdateOrderStatus(orderId string, statusReq *dto.UpdateOrderStatusRequest) error {
	order, err := s.repo.GetOrderSummaryByID(orderId)
	if err != nil {
		return err
	}

	if statusReq.Status == enums.OrderStatusCancelled {
		if order.Status != enums.OrderStatusPending &&
			order.Status != enums.OrderStatusConfirmed &&
			order.Status != enums.OrderStatusProcessing {
			return fmt.Errorf("order with status %s cannot be cancelled", order.Status)
		}
	}

	return s.repo.UpdateOrderStatus(orderId, statusReq.Status)
}
