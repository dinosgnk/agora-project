package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	"github.com/dinosgnk/agora-project/internal/pkg/rabbitmq"
	"github.com/dinosgnk/agora-project/internal/services/notification/consumer"
)

func main() {
	log := logger.NewLogger()

	rabbitClient, err := rabbitmq.NewRabbitMQClient(log)
	if err != nil {
		log.Error("Failed to connect to RabbitMQ", "error", err)
		os.Exit(1)
	}
	defer rabbitClient.Close()

	eventConsumer, err := consumer.NewEventConsumer(rabbitClient, log)
	if err != nil {
		log.Error("Failed to initialize event consumer", "error", err)
		os.Exit(1)
	}

	// Start consuming messages
	if err := eventConsumer.Start(); err != nil {
		log.Error("Failed to start event consumer", "error", err)
		os.Exit(1)
	}

	log.Info("Notification service started successfully")

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down notification service...")
}
