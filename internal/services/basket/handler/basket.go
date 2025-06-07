package handler

import (
	"encoding/json"
	"net/http"

	"agora/basket/model"
	"agora/basket/service"
)

type BasketHandler struct {
	service service.IBasketService
}

func NewBasketHandler(svc service.BasketService) *BasketHandler {
	return &BasketHandler{svc: svc}
}

func (h *BasketHandler) GetBasket(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		http.Error(w, "missing user_id", http.StatusBadRequest)
		return
	}

	basket, err := h.svc.GetBasket(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(basket)
}

func (h *BasketHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	if userId == "" {
		http.Error(w, "missing user_id", http.StatusBadRequest)
		return
	}

	var item model.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "invalid item", http.StatusBadRequest)
		return
	}

	if err := h.svc.AddItem(userId, item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *BasketHandler) ClearBasket(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	if userId == "" {
		http.Error(w, "missing user_id", http.StatusBadRequest)
		return
	}

	if err := h.svc.ClearBasket(userId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
