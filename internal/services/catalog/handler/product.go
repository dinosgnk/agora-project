package handler

import (
	"net/http"

	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	"github.com/dinosgnk/agora-project/internal/services/catalog/dto"
	"github.com/dinosgnk/agora-project/internal/services/catalog/model"
	"github.com/dinosgnk/agora-project/internal/services/catalog/service"
	"github.com/gin-gonic/gin"
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

func (h *ProductHandler) GetAllProducts(ctx *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		h.log.Error("Failed to get all products", "error", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductsByCategory(ctx *gin.Context) {
	category := ctx.Param("category")

	products, err := h.service.GetProductsByCategory(category)
	if err != nil {
		h.log.Error("Failed to get products by category", "category", category, "error", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductByCode(ctx *gin.Context) {
	productCode := ctx.Param("productCode")

	product, err := h.service.GetProductByCode(productCode)
	if err != nil {
		h.log.Error("Failed to get product by code", "product_code", productCode, "error", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var reqProduct dto.CreateProductRequest
	if err := ctx.ShouldBindJSON(&reqProduct); err != nil {
		h.log.Warn("Invalid request body for create product", "error", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := h.service.CreateProduct(&reqProduct)
	if err != nil {
		h.log.Error("Failed to create product", "product_code", reqProduct.ProductCode, "error", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdProduct)
}

func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	productCode := ctx.Param("productCode")

	var req dto.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.log.Warn("Invalid request body for update product", "product_code", productCode, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := dto.ProductResponse{
		ProductCode: updatedProduct.ProductCode,
		Name:        updatedProduct.Name,
		Category:    updatedProduct.Category,
		Description: updatedProduct.Description,
		Price:       updatedProduct.Price,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	productCode := c.Param("productCode")

	deleted, err := h.service.DeleteProduct(productCode)
	if err != nil {
		h.log.Error("Failed to delete product", "product_code", productCode, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !deleted {
		h.log.Warn("Product not found for deletion", "product_code", productCode)
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
