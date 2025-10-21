package main

import (
	"os"

	confighelper "github.com/dinosgnk/agora-project/internal/pkg/config"
	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	"github.com/dinosgnk/agora-project/internal/pkg/rabbitmq"
	"github.com/dinosgnk/agora-project/internal/pkg/server"
	"github.com/dinosgnk/agora-project/internal/services/order/config"
	"github.com/dinosgnk/agora-project/internal/services/order/handler"
	"github.com/dinosgnk/agora-project/internal/services/order/messaging"
	"github.com/dinosgnk/agora-project/internal/services/order/repository"
	"github.com/dinosgnk/agora-project/internal/services/order/service"
)

func main() {
	log := logger.NewLogger()
	cfg := confighelper.LoadConfig[config.AppConfig](log)

	rabbitClient, err := rabbitmq.NewRabbitMQClient(log)
	if err != nil {
		log.Error("Failed to connect to RabbitMQ", "error", err)
		os.Exit(1)
	}
	defer rabbitClient.Close()

	publisher, err := messaging.NewPublisher(rabbitClient)
	if err != nil {
		log.Error("Failed to initialize event publisher", "error", err)
		os.Exit(1)
	}

	orderRepository := repository.NewPostgresOrderRepository(log)
	orderService := service.NewOrderService(orderRepository, publisher)
	orderHandler := handler.NewOrderHandler(orderService, log)

	server := server.NewServer(cfg.Port, orderHandler, log, cfg.Service)
	if err := server.Run(); err != nil {
		os.Exit(1)
	}
}
