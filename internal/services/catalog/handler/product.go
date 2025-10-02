package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	"github.com/dinosgnk/agora-project/internal/services/catalog/dto"
	"github.com/dinosgnk/agora-project/internal/services/catalog/model"
	"github.com/dinosgnk/agora-project/internal/services/catalog/service"
)

type ProductHandler struct {
	service *service.ProductService
	log     logger.Logger
}

func NewProductHandler(s *service.ProductService, l logger.Logger) *ProductHandler {
	return &ProductHandler{
		service: s,
		log:     l,
	}
}

func (h *ProductHandler) RegisterRoutes(mux *http.ServeMux) http.Handler {
	mux.HandleFunc("GET /products", h.GetAllProducts)
	mux.HandleFunc("GET /products/category/{category}", h.GetProductsByCategory)
	mux.HandleFunc("GET /products/{productCode}", h.GetProductByCode)
	mux.HandleFunc("POST /products", h.CreateProduct)
	mux.HandleFunc("PUT /products/{productCode}", h.UpdateProduct)
	mux.HandleFunc("DELETE /products/{productCode}", h.DeleteProduct)

	return mux
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		h.log.Error("Failed to get all products", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	category := r.PathValue("category")

	products, err := h.service.GetProductsByCategory(category)
	if err != nil {
		h.log.Error("Failed to get products by category", "category", category, "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProductByCode(w http.ResponseWriter, r *http.Request) {
	productCode := r.PathValue("productCode")

	product, err := h.service.GetProductByCode(productCode)
	if err != nil {
		h.log.Error("Failed to get product by code", "product_code", productCode, "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var reqProduct dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&reqProduct); err != nil {
		h.log.Warn("Invalid request body for create product", "error", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdProduct, err := h.service.CreateProduct(&reqProduct)
	if err != nil {
		h.log.Error("Failed to create product", "product_code", reqProduct.ProductCode, "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProduct)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productCode := r.PathValue("productCode")

	var req dto.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Warn("Invalid request body for update product", "product_code", productCode, "error", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product := &model.Product{
		ProductCode: req.ProductCode,
		Name:        req.Name,
		Category:    req.Category,
		Description: req.Description,
		Price:       req.Price,
	}

	updatedProduct, err := h.service.UpdateProduct(productCode, product)
	if err != nil {
		h.log.Error("Failed to update product", "product_code", productCode, "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.ProductResponse{
		ProductCode: updatedProduct.ProductCode,
		Name:        updatedProduct.Name,
		Category:    updatedProduct.Category,
		Description: updatedProduct.Description,
		Price:       updatedProduct.Price,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productCode := r.PathValue("productCode")

	deleted, err := h.service.DeleteProduct(productCode)
	if err != nil {
		h.log.Error("Failed to delete product", "product_code", productCode, "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !deleted {
		h.log.Warn("Product not found for deletion", "product_code", productCode)
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
