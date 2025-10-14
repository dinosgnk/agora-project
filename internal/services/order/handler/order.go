package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	"github.com/dinosgnk/agora-project/internal/services/order/dto"
	"github.com/dinosgnk/agora-project/internal/services/order/service"
)

type OrderHandler struct {
	service service.IOrderService
	log     logger.Logger
}

func NewOrderHandler(s service.IOrderService, l logger.Logger) *OrderHandler {
	return &OrderHandler{
		service: s,
		log:     l,
	}
}

func (h *OrderHandler) RegisterRoutes(mux *http.ServeMux) http.Handler {
	mux.HandleFunc("POST /orders", h.CreateOrder)
	mux.HandleFunc("GET /orders/summary", h.GetAllOrderSummaries)
	mux.HandleFunc("GET /orders", h.GetAllOrders)
	mux.HandleFunc("GET /orders/user/{userId}/summary", h.GetAllOrderSummariesByUserID)
	mux.HandleFunc("GET /orders/user/{userId}", h.GetAllOrdersByUserID)
	mux.HandleFunc("GET /orders/order/{orderId}/summary", h.GetOrderSummaryByID)
	mux.HandleFunc("GET /orders/order/{orderId}/products", h.GetProductsByOrderID)
	mux.HandleFunc("GET /orders/order/{orderId}", h.GetOrderByID)
	mux.HandleFunc("PUT /orders/order/{orderId}/status", h.UpdateOrderStatus)

	return mux
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderReq dto.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&orderReq); err != nil {
		h.log.Warn("Invalid request body for create order", "error", err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdOrder, err := h.service.CreateOrder(&orderReq)
	if err != nil {
		h.log.Error("Failed to create order", "user_id", orderReq.UserID, "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdOrder)
}

func (h *OrderHandler) GetAllOrderSummaries(w http.ResponseWriter, r *http.Request) {
	orders, err := h.service.GetAllOrderSummaries()
	if err != nil {
		h.log.Error("Failed to get all orders summary", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.service.GetAllOrders()
	if err != nil {
		h.log.Error("Failed to get all orders", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) GetOrderSummaryByID(w http.ResponseWriter, r *http.Request) {
	orderId := r.PathValue("orderId")

	order, err := h.service.GetOrderSummaryByID(orderId)
	if err != nil {
		h.log.Error("Failed to get order by id", "order_id", orderId, "error", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	orderId := r.PathValue("orderId")

	order, err := h.service.GetOrderByID(orderId)
	if err != nil {
		h.log.Error("Failed to get order by id", "order_id", orderId, "error", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetAllOrderSummariesByUserID(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")

	orders, err := h.service.GetAllOrderSummariesByUserID(userId)
	if err != nil {
		h.log.Error("Failed to get all orders summary by user id", "user_id", userId, "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) GetAllOrdersByUserID(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")

	orders, err := h.service.GetAllOrdersByUserID(userId)
	if err != nil {
		h.log.Error("Failed to get orders by user id", "user_id", userId, "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) GetProductsByOrderID(w http.ResponseWriter, r *http.Request) {
	orderId := r.PathValue("orderId")
	products, err := h.service.GetProductsByOrderID(orderId)
	if err != nil {
		h.log.Error("Failed to get products by order id", "order_id", orderId, "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (h *OrderHandler) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	orderId := r.PathValue("orderId")

	var statusReq dto.UpdateOrderStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&statusReq); err != nil {
		h.log.Warn("Invalid request body for update order status", "error", err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateOrderStatus(orderId, &statusReq); err != nil {
		h.log.Error("Failed to update order status", "order_id", orderId, "status", statusReq.Status, "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
