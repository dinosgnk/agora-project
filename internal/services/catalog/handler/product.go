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
	var products []*model.Product
	products, err := h.service.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductsByCategory(ctx *gin.Context) {
	category := ctx.Param("category")
	var products []*model.Product
	products, err := h.service.GetProductsByCategory(category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductByCode(ctx *gin.Context) {
	productCode := ctx.Param("productCode")
	product, err := h.service.GetProductByCode(productCode)
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
		ProductCode: req.ProductCode,
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
		ProductCode: createdProduct.ProductCode,
		Name:        createdProduct.Name,
		Category:    createdProduct.Category,
		Description: createdProduct.Description,
		Price:       createdProduct.Price,
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	productCode := ctx.Param("productCode")
	var req dto.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
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
	_, err := h.service.DeleteProduct(productCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
