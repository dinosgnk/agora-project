package main

import (
	"os"

	confighelper "github.com/dinosgnk/agora-project/internal/pkg/config"
	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	"github.com/dinosgnk/agora-project/internal/pkg/server"
	"github.com/dinosgnk/agora-project/internal/services/cart/config"
	"github.com/dinosgnk/agora-project/internal/services/cart/handler"
	"github.com/dinosgnk/agora-project/internal/services/cart/repository"
	"github.com/dinosgnk/agora-project/internal/services/cart/service"
)

func main() {
	log := logger.NewLogger()
	cfg := confighelper.LoadConfig[config.AppConfig](log)

	cartRepository := repository.NewInMemoryRepository()
	cartService := service.NewCartService(cartRepository)
	cartHandler := handler.NewCartHandler(cartService, log)

	server := server.NewServer(cfg.Port, cartHandler, log)
	if err := server.Run(); err != nil {
		os.Exit(1)
	}
}
