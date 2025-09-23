package main

import (
	confighelper "github.com/dinosgnk/agora-project/internal/pkg/config"
	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	middleware "github.com/dinosgnk/agora-project/internal/pkg/middleware/logging"
	"github.com/dinosgnk/agora-project/internal/services/catalog/config"
	"github.com/dinosgnk/agora-project/internal/services/catalog/handler"
	"github.com/dinosgnk/agora-project/internal/services/catalog/metrics"
	"github.com/dinosgnk/agora-project/internal/services/catalog/repository"
	"github.com/dinosgnk/agora-project/internal/services/catalog/service"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log := logger.NewLogger()

	cfg := confighelper.LoadConfig[config.AppConfig](log)

	log.Info("Starting catalog service", "environment", cfg.Environment, "port", cfg.Port)

	productRepository := repository.NewPostgresProductRepository(log)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService, log)

	router := gin.New()
	router.Use(middleware.RequestLoggingMiddleware(log))
	router.Use(metrics.PrometheusMiddleware())

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.GET("/health", func(ctx *gin.Context) {
		ctx.String(200, "Catalog service is healthy")
	})
	router.GET("/products", productHandler.GetAllProducts)
	router.GET("/products/category/:category", productHandler.GetProductsByCategory)
	router.GET("/products/:productCode", productHandler.GetProductByCode)
	router.POST("/products/:productCode", productHandler.CreateProduct)
	router.PUT("/products/:productCode", productHandler.UpdateProduct)
	router.DELETE("/products/:productCode", productHandler.DeleteProduct)

	log.Info("Catalog service listening on port", "port", cfg.Port)
	router.Run(":" + cfg.Port)
}
