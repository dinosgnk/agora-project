package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dinosgnk/agora-project/internal/services/cart/dto"
	"github.com/dinosgnk/agora-project/internal/services/cart/model"
	"github.com/dinosgnk/agora-project/internal/services/cart/service"
)

type CartHandler struct {
	service service.ICartService
}

func NewCartHandler(service service.ICartService) *CartHandler {
	return &CartHandler{service: service}
}

func (ch *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		http.Error(w, "Missing userId", http.StatusBadRequest)
		return
	}

	basket, err := ch.service.GetCartByUserId(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(basket)
}

func (ch *CartHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	var req dto.AddItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.UserId == "" || req.Item.ProductCode == "" {
		http.Error(w, "Missing user_id or product_code", http.StatusBadRequest)
		return
	}

	itemToAdd := model.Item{
		ProductCode: req.Item.ProductCode,
		Name:        req.Item.Name,
		Quantity:    req.Item.Quantity,
		Price:       req.Item.Price,
	}

	if err := ch.service.AddItem(req.UserId, &itemToAdd); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ch *CartHandler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	var req dto.RemoveItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.UserId == "" || req.ProductCode == "" {
		http.Error(w, "Missing user_id or product_code", http.StatusBadRequest)
		return
	}

	if err := ch.service.RemoveItem(req.UserId, req.ProductCode); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ch *CartHandler) UpdateCart(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if len(req.Items) == 0 {
		w.WriteHeader(http.StatusNoContent)
	}

	if err := ch.service.UpdateCart(req.UserId, req.Items); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (ch *CartHandler) ClearCart(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		http.Error(w, "missing userId", http.StatusBadRequest)
		return
	}

	if err := ch.service.ClearCart(userId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
