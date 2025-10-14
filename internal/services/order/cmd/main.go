package main

import (
	"os"

	confighelper "github.com/dinosgnk/agora-project/internal/pkg/config"
	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	"github.com/dinosgnk/agora-project/internal/pkg/server"
	"github.com/dinosgnk/agora-project/internal/services/order/config"
	"github.com/dinosgnk/agora-project/internal/services/order/handler"
	"github.com/dinosgnk/agora-project/internal/services/order/repository"
	"github.com/dinosgnk/agora-project/internal/services/order/service"
)

func main() {
	log := logger.NewLogger()
	cfg := confighelper.LoadConfig[config.AppConfig](log)

	orderRepository := repository.NewPostgresOrderRepository(log)
	orderService := service.NewOrderService(orderRepository)
	orderHandler := handler.NewOrderHandler(orderService, log)

	server := server.NewServer(cfg.Port, orderHandler, log, cfg.Service)
	if err := server.Run(); err != nil {
		os.Exit(1)
	}
}
