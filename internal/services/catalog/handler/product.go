package handler

import (
	"net/http"

	"github.com/dinosgnk/agora-project/internal/services/catalog/dto"
	"github.com/dinosgnk/agora-project/internal/services/catalog/model"
	"github.com/dinosgnk/agora-project/internal/services/catalog/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) GetAllProducts(ctx *gin.Context) {
	var products *[]model.Product
	products, err := h.service.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductsByCategory(ctx *gin.Context) {
	category := ctx.Param("category")
	var products *[]model.Product
	products, err := h.service.GetProductsByCategory(category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := h.service.GetProductById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var req dto.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := model.Product{
		Name:        req.Name,
		Category:    req.Category,
		Description: req.Description,
		Price:       req.Price,
	}

	createdProduct, err := h.service.CreateProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := dto.ProductResponse{
		ProductId:   createdProduct.ProductId,
		Name:        createdProduct.Name,
		Category:    createdProduct.Category,
		Description: createdProduct.Description,
		Price:       createdProduct.Price,
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var req dto.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := model.Product{
		ProductId:   id,
		Name:        req.Name,
		Category:    req.Category,
		Description: req.Description,
		Price:       req.Price,
	}

	updatedProduct, err := h.service.UpdateProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := dto.ProductResponse{
		ProductId:   updatedProduct.ProductId,
		Name:        updatedProduct.Name,
		Category:    updatedProduct.Category,
		Description: updatedProduct.Description,
		Price:       updatedProduct.Price,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	_, err := h.service.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
