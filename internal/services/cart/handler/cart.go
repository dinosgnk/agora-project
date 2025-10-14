package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	"github.com/dinosgnk/agora-project/internal/services/cart/dto"
	"github.com/dinosgnk/agora-project/internal/services/cart/service"
)

type CartHandler struct {
	service service.ICartService
	log     logger.Logger
}

func NewCartHandler(s service.ICartService, l logger.Logger) *CartHandler {
	return &CartHandler{
		service: s,
		log:     l,
	}
}

func (h *CartHandler) RegisterRoutes(mux *http.ServeMux) http.Handler {
	mux.HandleFunc("/cart/{userId}", h.GetCart)
	mux.HandleFunc("/cart/item/add/{userId}", h.AddItem)
	mux.HandleFunc("/cart/item/delete/{userId}", h.RemoveItem)
	mux.HandleFunc("/cart/update/{userId}", h.UpdateCart)
	mux.HandleFunc("/cart/clear/{userId}", h.ClearCart)
	return mux
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	basket, err := h.service.GetCartByUserId(userId)
	if err != nil {
		h.log.Error("Failed to get cart", "user_id", userId, "error", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(basket)
}

func (h *CartHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")

	var req dto.AddItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Warn("Invalid request body for add item", "error", err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	itemToAdd := &dto.Item{
		ProductCode: req.Item.ProductCode,
		Name:        req.Item.Name,
		Quantity:    req.Item.Quantity,
		Price:       req.Item.Price,
	}

	if err := h.service.AddItem(userId, itemToAdd); err != nil {
		h.log.Error("Failed to add item to cart", "user_id", userId, "product_code", req.Item.ProductCode, "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *CartHandler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")

	var req dto.RemoveItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Warn("Invalid request body for remove item", "error", err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.RemoveItem(userId, req.ProductCode); err != nil {
		h.log.Error("Failed to remove item from cart", "user_id", userId, "product_code", req.ProductCode, "error", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *CartHandler) UpdateCart(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")

	var req dto.UpdateCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Warn("Invalid request body for update cart", "error", err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateCart(userId, req.Items); err != nil {
		h.log.Error("Failed to update cart", "user_id", userId, "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *CartHandler) ClearCart(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")

	if err := h.service.ClearCart(userId); err != nil {
		h.log.Error("Failed to clear cart", "user_id", userId, "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
