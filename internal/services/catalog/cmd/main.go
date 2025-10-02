package main

import (
	"os"

	confighelper "github.com/dinosgnk/agora-project/internal/pkg/config"
	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	"github.com/dinosgnk/agora-project/internal/pkg/server"
	"github.com/dinosgnk/agora-project/internal/services/catalog/config"
	"github.com/dinosgnk/agora-project/internal/services/catalog/handler"
	"github.com/dinosgnk/agora-project/internal/services/catalog/repository"
	"github.com/dinosgnk/agora-project/internal/services/catalog/service"
)

func main() {
	log := logger.NewLogger()
	cfg := confighelper.LoadConfig[config.AppConfig](log)

	productRepository := repository.NewPostgresProductRepository(log)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService, log)

	server := server.NewServer(cfg.Port, productHandler, log)
	if err := server.Run(); err != nil {
		os.Exit(1)
	}
}
