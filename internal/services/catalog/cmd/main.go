package main

import (
	"github.com/dinosgnk/agora-project/internal/services/catalog/handler"
	"github.com/dinosgnk/agora-project/internal/services/catalog/repository"
	"github.com/dinosgnk/agora-project/internal/services/catalog/service"
	"github.com/gin-gonic/gin"
)

func main() {
	productRepository := repository.NewPostgresProductRepository()
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	router := gin.Default()
	router.GET("/health", func(ctx *gin.Context) {
		ctx.String(200, "Catalog service is healthy")
	})
	router.GET("/products", productHandler.GetAllProducts)
	router.GET("/products/category/:category", productHandler.GetProductsByCategory)
	router.GET("/products/:id", productHandler.GetProductById)
	router.POST("/products/:id", productHandler.CreateProduct)
	router.PUT("/products/:id", productHandler.UpdateProduct)
	router.DELETE("/products/:id", productHandler.DeleteProduct)

	router.Run(":8080") // Listen and serve on 0.0.0.0:8080
}
