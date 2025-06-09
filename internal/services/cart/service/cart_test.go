package service

import (
	"testing"

	"github.com/dinosgnk/agora-project/internal/services/cart/model"
	"github.com/dinosgnk/agora-project/internal/services/cart/repository"
)

func TestGetCartByUserIdSuccessfully(t *testing.T) {
	repo := repository.NewMockCartRepository()
	svc := NewCartService(repo)

	userId := "10"
	itemToAdd := model.Item{
		ProductId: "1099",
		Name:      "XYZ Product Name",
		Price:     10.0,
		Quantity:  1,
	}
	err := svc.AddItem(userId, &itemToAdd)
	if err != nil {
		t.Fatalf("Expected no error while adding item to cart, got %v", err)
	}

	cart, err := svc.GetCartByUserId(userId)
	if err != nil {
		t.Fatalf("Expected cart, got error %v", err)
	}
	if len(cart.Items) != 1 {
		t.Fatalf("Expected 1 item in the cart, got %d", len(cart.Items))
	}
}

func TestAddItemToCartSuccessfully(t *testing.T) {
	repo := repository.NewMockCartRepository()
	svc := NewCartService(repo)

	userId := "10"
	itemToAdd := model.Item{
		ProductId: "1099",
		Name:      "XYZ Product Name",
		Price:     10.0,
		Quantity:  1,
	}

	err := svc.AddItem(userId, &itemToAdd)
	if err != nil {
		t.Fatalf("Expected no error while adding item to cart, got %v", err)
	}

	cart, err := repo.GetCartByUserId(userId)
	if err != nil {
		t.Fatalf("Expected cart, got error %v", err)
	}

	if len(cart.Items) != 1 {
		t.Fatalf("Expected 1 item, got %d", len(cart.Items))
	}
}

func TestRemoveItemFromCartSuccessfully(t *testing.T) {
	repo := repository.NewMockCartRepository()
	svc := NewCartService(repo)

	userID := "user123"
	itemToAdd := model.Item{ProductId: "p1", Name: "Product", Price: 10.0, Quantity: 1}
	_ = svc.AddItem(userID, &itemToAdd)

	err := svc.RemoveItem(userID, "p1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	cart, _ := repo.GetCartByUserId(userID)
	if len(cart.Items) != 0 {
		t.Fatalf("Expected item to be removed, got %d items", len(cart.Items))
	}
}

func TestClearCart(t *testing.T) {
	repo := repository.NewMockCartRepository()
	svc := NewCartService(repo)

	userId := "10"
	itemToAdd := model.Item{
		ProductId: "1099",
		Name:      "XYZ Product Name",
		Price:     10.0,
		Quantity:  1,
	}
	err := svc.AddItem(userId, &itemToAdd)
	if err != nil {
		t.Fatalf("Expected no error while adding item to cart, got %v", err)
	}

	err = svc.ClearCart(userId)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	_, err = repo.GetCartByUserId(userId)
	if err == nil {
		t.Fatalf("Expected error after clearing cart, got none")
	}
}
