package main

import (
	"github.com/dinosgnk/agora-project/internal/services/catalog/handler"
	"github.com/dinosgnk/agora-project/internal/services/catalog/metrics"
	"github.com/dinosgnk/agora-project/internal/services/catalog/repository"
	"github.com/dinosgnk/agora-project/internal/services/catalog/service"
	"github.com/dinosgnk/agora-project/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	cfg := config.LoadConfig()

	productRepository := repository.NewPostgresProductRepository()
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	router := gin.Default()

	router.Use(metrics.PrometheusMiddleware())

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.GET("/health", func(ctx *gin.Context) {
		ctx.String(200, "Catalog service is healthy")
	})
	router.GET("/products", productHandler.GetAllProducts)
	router.GET("/products/category/:category", productHandler.GetProductsByCategory)
	router.GET("/products/:id", productHandler.GetProductById)
	router.POST("/products/:id", productHandler.CreateProduct)
	router.PUT("/products/:id", productHandler.UpdateProduct)
	router.DELETE("/products/:id", productHandler.DeleteProduct)

	router.Run(":" + cfg.Port)
}
