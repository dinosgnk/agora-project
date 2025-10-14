package service

import (
	"testing"

	"github.com/dinosgnk/agora-project/internal/services/order/dto"
	"github.com/dinosgnk/agora-project/internal/services/order/enums"
	"github.com/dinosgnk/agora-project/internal/services/order/repository"
)

func TestCreateOrderSuccessfully(t *testing.T) {
	repo := repository.NewMockOrderRepository()
	svc := NewOrderService(repo)

	orderReq := &dto.CreateOrderRequest{
		UserID: "user123",
		Products: []*dto.OrderedProduct{
			{
				ProductCode: "P1",
				ProductName: "Product 1",
				Quantity:    2,
				Price:       10.99,
			},
			{
				ProductCode: "P2",
				ProductName: "Product 2",
				Quantity:    1,
				Price:       25.50,
			},
		},
		ShippingAddress: "Address 123",
		PaymentMethod:   "crypto",
	}

	orderResp, err := svc.CreateOrder(orderReq)
	if err != nil {
		t.Fatalf("Expected no error while creating order, got %v", err)
	}

	if orderResp.OrderID == "" {
		t.Fatal("Expected order ID to be set")
	}

	if orderResp.UserID != orderReq.UserID {
		t.Fatalf("Expected user ID %s, got %s", orderReq.UserID, orderResp.UserID)
	}

	if orderResp.Status != enums.OrderStatusPending {
		t.Fatalf("Expected status %s, got %s", enums.OrderStatusPending, orderResp.Status)
	}

	expectedTotal := 2*10.99 + 1*25.50
	floatTolerance := 0.0001
	if diff := orderResp.TotalAmount - expectedTotal; diff > floatTolerance || diff < -floatTolerance {
		t.Fatalf("Expected total amount %.2f, got %.2f", expectedTotal, orderResp.TotalAmount)
	}

	if orderResp.ShippingAddress != orderReq.ShippingAddress {
		t.Fatalf("Expected shipping address %s, got %s", orderReq.ShippingAddress, orderResp.ShippingAddress)
	}

	if orderResp.PaymentMethod != orderReq.PaymentMethod {
		t.Fatalf("Expected payment method %s, got %s", orderReq.PaymentMethod, orderResp.PaymentMethod)
	}

	if len(orderResp.Products) != len(orderReq.Products) {
		t.Fatalf("Expected %d products, got %d", len(orderReq.Products), len(orderResp.Products))
	}
}

func TestGetAllOrderSummariesSuccessfully(t *testing.T) {
	repo := repository.NewMockOrderRepository()
	svc := NewOrderService(repo)

	orderReq1 := &dto.CreateOrderRequest{
		UserID: "user123",
		Products: []*dto.OrderedProduct{
			{ProductCode: "P1", ProductName: "Product 1", Quantity: 1, Price: 10.00},
		},
		ShippingAddress: "Address 123",
		PaymentMethod:   "crypto",
	}
	_, _ = svc.CreateOrder(orderReq1)

	orderReq2 := &dto.CreateOrderRequest{
		UserID: "user456",
		Products: []*dto.OrderedProduct{
			{ProductCode: "P2", ProductName: "Product 2", Quantity: 2, Price: 20.00},
		},
		ShippingAddress: "Address 2",
		PaymentMethod:   "crypto",
	}
	_, _ = svc.CreateOrder(orderReq2)

	summaries, err := svc.GetAllOrderSummaries()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(summaries) != 2 {
		t.Fatalf("Expected 2 order summaries, got %d", len(summaries))
	}
}

func TestGetAllOrdersSuccessfully(t *testing.T) {
	repo := repository.NewMockOrderRepository()
	svc := NewOrderService(repo)

	orderReq := &dto.CreateOrderRequest{
		UserID: "user123",
		Products: []*dto.OrderedProduct{
			{ProductCode: "P1", ProductName: "Product 1", Quantity: 1, Price: 10.00},
			{ProductCode: "P2", ProductName: "Product 2", Quantity: 3, Price: 5.00},
		},
		ShippingAddress: "Address 123",
		PaymentMethod:   "crypto",
	}
	_, _ = svc.CreateOrder(orderReq)

	orders, err := svc.GetAllOrders()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(orders) != 1 {
		t.Fatalf("Expected 1 order, got %d", len(orders))
	}

	if orders[0].UserID != "user123" {
		t.Fatalf("Expected user ID user123, got %s", orders[0].UserID)
	}
}

func TestGetOrderSummaryByIDSuccessfully(t *testing.T) {
	repo := repository.NewMockOrderRepository()
	svc := NewOrderService(repo)

	orderReq := &dto.CreateOrderRequest{
		UserID: "user123",
		Products: []*dto.OrderedProduct{
			{ProductCode: "P1", ProductName: "Product 1", Quantity: 1, Price: 10.00},
		},
		ShippingAddress: "Address 123",
		PaymentMethod:   "crypto",
	}
	createdOrder, _ := svc.CreateOrder(orderReq)

	summary, err := svc.GetOrderSummaryByID(createdOrder.OrderID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if summary.OrderID != createdOrder.OrderID {
		t.Fatalf("Expected order ID %s, got %s", createdOrder.OrderID, summary.OrderID)
	}

	if summary.UserID != "user123" {
		t.Fatalf("Expected user ID user123, got %s", summary.UserID)
	}
}

func TestGetOrderByIDSuccessfully(t *testing.T) {
	repo := repository.NewMockOrderRepository()
	svc := NewOrderService(repo)

	orderReq := &dto.CreateOrderRequest{
		UserID: "user123",
		Products: []*dto.OrderedProduct{
			{ProductCode: "P1", ProductName: "Product 1", Quantity: 2, Price: 15.00},
			{ProductCode: "P2", ProductName: "Product 2", Quantity: 1, Price: 30.00},
		},
		ShippingAddress: "Address 123",
		PaymentMethod:   "crypto",
	}
	createdOrder, _ := svc.CreateOrder(orderReq)

	order, err := svc.GetOrderByID(createdOrder.OrderID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if order.OrderID != createdOrder.OrderID {
		t.Fatalf("Expected order ID %s, got %s", createdOrder.OrderID, order.OrderID)
	}

	if len(order.Products) != 2 {
		t.Fatalf("Expected 2 products, got %d", len(order.Products))
	}
}

func TestGetAllOrderSummariesByUserIDSuccessfully(t *testing.T) {
	repo := repository.NewMockOrderRepository()
	svc := NewOrderService(repo)

	userId := "user123"

	// Create multiple orders for the same user
	for i := 0; i < 3; i++ {
		orderReq := &dto.CreateOrderRequest{
			UserID: userId,
			Products: []*dto.OrderedProduct{
				{ProductCode: "P1", ProductName: "Product", Quantity: 1, Price: 10.00},
			},
			ShippingAddress: "Address 123",
			PaymentMethod:   "crypto",
		}
		_, _ = svc.CreateOrder(orderReq)
	}

	// Create order for different user
	otherOrderReq := &dto.CreateOrderRequest{
		UserID: "user456",
		Products: []*dto.OrderedProduct{
			{ProductCode: "P2", ProductName: "Product", Quantity: 1, Price: 20.00},
		},
		ShippingAddress: "Other Address",
		PaymentMethod:   "paypal",
	}
	_, _ = svc.CreateOrder(otherOrderReq)

	// Get orders for user123
	summaries, err := svc.GetAllOrderSummariesByUserID(userId)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(summaries) != 3 {
		t.Fatalf("Expected 3 order summaries, got %d", len(summaries))
	}

	for _, summary := range summaries {
		if summary.UserID != userId {
			t.Fatalf("Expected user ID %s, got %s", userId, summary.UserID)
		}
	}
}

func TestGetAllOrdersByUserIDSuccessfully(t *testing.T) {
	repo := repository.NewMockOrderRepository()
	svc := NewOrderService(repo)

	userId := "user123"

	orderReq := &dto.CreateOrderRequest{
		UserID: userId,
		Products: []*dto.OrderedProduct{
			{ProductCode: "P1", ProductName: "Product 1", Quantity: 1, Price: 10.00},
			{ProductCode: "P2", ProductName: "Product 2", Quantity: 2, Price: 15.00},
		},
		ShippingAddress: "Address 123",
		PaymentMethod:   "crypto",
	}
	_, _ = svc.CreateOrder(orderReq)

	orders, err := svc.GetAllOrdersByUserID(userId)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(orders) != 1 {
		t.Fatalf("Expected 1 order, got %d", len(orders))
	}

	if orders[0].UserID != userId {
		t.Fatalf("Expected user ID %s, got %s", userId, orders[0].UserID)
	}
}

func TestGetProductsByOrderIDSuccessfully(t *testing.T) {
	repo := repository.NewMockOrderRepository()
	svc := NewOrderService(repo)

	orderReq := &dto.CreateOrderRequest{
		UserID: "user123",
		Products: []*dto.OrderedProduct{
			{ProductCode: "P1", ProductName: "Product 1", Quantity: 1, Price: 10.00},
			{ProductCode: "P2", ProductName: "Product 2", Quantity: 3, Price: 5.00},
		},
		ShippingAddress: "Address 123",
		PaymentMethod:   "crypto",
	}
	createdOrder, _ := svc.CreateOrder(orderReq)

	products, err := svc.GetProductsByOrderID(createdOrder.OrderID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(products) != 2 {
		t.Fatalf("Expected 2 products, got %d", len(products))
	}

	if products[0].ProductCode != "P1" {
		t.Fatalf("Expected product code P1, got %s", products[0].ProductCode)
	}

	if products[1].ProductCode != "P2" {
		t.Fatalf("Expected product code P2, got %s", products[0].ProductCode)
	}

	if products[1].Quantity != 3 {
		t.Fatalf("Expected quantity 3, got %d", products[1].Quantity)
	}
}

func TestUpdateOrderStatusSuccessfully(t *testing.T) {
	repo := repository.NewMockOrderRepository()
	svc := NewOrderService(repo)

	orderReq := &dto.CreateOrderRequest{
		UserID: "user123",
		Products: []*dto.OrderedProduct{
			{ProductCode: "P1", ProductName: "Product", Quantity: 1, Price: 10.00},
		},
		ShippingAddress: "Address 123",
		PaymentMethod:   "crypto",
	}
	createdOrder, _ := svc.CreateOrder(orderReq)

	statusReq := &dto.UpdateOrderStatusRequest{
		Status: enums.OrderStatusConfirmed,
	}
	err := svc.UpdateOrderStatus(createdOrder.OrderID, statusReq)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	order, _ := svc.GetOrderSummaryByID(createdOrder.OrderID)
	if order.Status != enums.OrderStatusConfirmed {
		t.Fatalf("Expected status %s, got %s", enums.OrderStatusConfirmed, order.Status)
	}
}

func TestCancelOrderSuccessfully(t *testing.T) {
	repo := repository.NewMockOrderRepository()
	svc := NewOrderService(repo)

	orderReq := &dto.CreateOrderRequest{
		UserID: "user123",
		Products: []*dto.OrderedProduct{
			{ProductCode: "P1", ProductName: "Product", Quantity: 1, Price: 10.00},
		},
		ShippingAddress: "Address 123",
		PaymentMethod:   "crypto",
	}
	createdOrder, _ := svc.CreateOrder(orderReq)

	statusReq := &dto.UpdateOrderStatusRequest{
		Status: enums.OrderStatusCancelled,
	}
	err := svc.UpdateOrderStatus(createdOrder.OrderID, statusReq)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify order was cancelled
	order, _ := svc.GetOrderSummaryByID(createdOrder.OrderID)
	if order.Status != enums.OrderStatusCancelled {
		t.Fatalf("Expected status %s, got %s", enums.OrderStatusCancelled, order.Status)
	}
}

func TestCancelOrderFromConfirmedStatus(t *testing.T) {
	repo := repository.NewMockOrderRepository()
	svc := NewOrderService(repo)

	orderReq := &dto.CreateOrderRequest{
		UserID: "user123",
		Products: []*dto.OrderedProduct{
			{ProductCode: "P1", ProductName: "Product", Quantity: 1, Price: 10.00},
		},
		ShippingAddress: "Address 123",
		PaymentMethod:   "crypto",
	}
	createdOrder, _ := svc.CreateOrder(orderReq)

	statusReq := &dto.UpdateOrderStatusRequest{
		Status: enums.OrderStatusConfirmed,
	}
	_ = svc.UpdateOrderStatus(createdOrder.OrderID, statusReq)

	cancelReq := &dto.UpdateOrderStatusRequest{
		Status: enums.OrderStatusCancelled,
	}
	err := svc.UpdateOrderStatus(createdOrder.OrderID, cancelReq)
	if err != nil {
		t.Fatalf("Expected no error cancelling confirmed order, got %v", err)
	}

	order, _ := svc.GetOrderSummaryByID(createdOrder.OrderID)
	if order.Status != enums.OrderStatusCancelled {
		t.Fatalf("Expected status %s, got %s", enums.OrderStatusCancelled, order.Status)
	}
}

func TestCancelOrderWithInvalidStatus(t *testing.T) {
	repo := repository.NewMockOrderRepository()
	svc := NewOrderService(repo)

	orderReq := &dto.CreateOrderRequest{
		UserID: "user123",
		Products: []*dto.OrderedProduct{
			{ProductCode: "P1", ProductName: "Product", Quantity: 1, Price: 10.00},
		},
		ShippingAddress: "Address 123",
		PaymentMethod:   "crypto",
	}
	createdOrder, _ := svc.CreateOrder(orderReq)

	statusReq := &dto.UpdateOrderStatusRequest{
		Status: enums.OrderStatusShipped,
	}
	_ = svc.UpdateOrderStatus(createdOrder.OrderID, statusReq)

	cancelReq := &dto.UpdateOrderStatusRequest{
		Status: enums.OrderStatusCancelled,
	}
	err := svc.UpdateOrderStatus(createdOrder.OrderID, cancelReq)
	if err == nil {
		t.Fatal("Expected error when cancelling shipped order, got none")
	}
}

func TestCreateOrderCalculatesTotalCorrectly(t *testing.T) {
	repo := repository.NewMockOrderRepository()
	svc := NewOrderService(repo)

	testCases := []struct {
		name     string
		products []*dto.OrderedProduct
		expected float64
	}{
		{
			name: "Single product",
			products: []*dto.OrderedProduct{
				{ProductCode: "P1", ProductName: "Product 1", Quantity: 3, Price: 10.50},
			},
			expected: 31.50,
		},
		{
			name: "Multiple products",
			products: []*dto.OrderedProduct{
				{ProductCode: "P1", ProductName: "Product 1", Quantity: 2, Price: 10.00},
				{ProductCode: "P2", ProductName: "Product 2", Quantity: 5, Price: 3.50},
			},
			expected: 37.50,
		},
		{
			name: "Decimal precision",
			products: []*dto.OrderedProduct{
				{ProductCode: "P1", ProductName: "Product 1", Quantity: 3, Price: 9.99},
			},
			expected: 29.97,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			orderReq := &dto.CreateOrderRequest{
				UserID:          "user123",
				Products:        tc.products,
				ShippingAddress: "Address 123",
				PaymentMethod:   "crypto",
			}

			order, err := svc.CreateOrder(orderReq)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if order.TotalAmount != tc.expected {
				t.Fatalf("Expected total %.2f, got %.2f", tc.expected, order.TotalAmount)
			}
		})
	}
}
