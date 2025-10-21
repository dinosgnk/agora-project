package consumer

import (
	"encoding/json"
	"fmt"

	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	"github.com/dinosgnk/agora-project/internal/pkg/rabbitmq"
)

const (
	ordersExchange     = "orders"
	notificationsQueue = "notifications"
)

type OrderEvent struct {
	EventID   string `json:"event_id"`
	EventType string `json:"event_type"`
	Timestamp string `json:"timestamp"`
	OrderID   string `json:"order_id"`
	UserID    string `json:"user_id"`
}

type EventConsumer struct {
	client *rabbitmq.RabbitMQClient
	log    logger.Logger
}

func NewEventConsumer(client *rabbitmq.RabbitMQClient, log logger.Logger) (*EventConsumer, error) {
	if err := client.DeclareExchange(ordersExchange, "topic"); err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	if err := client.DeclareQueue(notificationsQueue); err != nil {
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	if err := client.BindQueue(notificationsQueue, ordersExchange, "order.*"); err != nil {
		return nil, fmt.Errorf("failed to bind queue: %w", err)
	}

	return &EventConsumer{
		client: client,
		log:    log,
	}, nil
}

func (c *EventConsumer) Start() error {
	c.log.Info("Starting event consumer", "queue", notificationsQueue)

	return c.client.Consume(notificationsQueue, c.handleMessage)
}

func (c *EventConsumer) handleMessage(body []byte) error {
	var event OrderEvent
	if err := json.Unmarshal(body, &event); err != nil {
		c.log.Error("Failed to unmarshal event", "error", err)
		return err
	}

	c.log.Info("Received order event",
		"event_id", event.EventID,
		"event_type", event.EventType,
		"order_id", event.OrderID,
		"user_id", event.UserID,
		"timestamp", event.Timestamp,
	)

	// TODO: Implement actual notification logic here
	c.log.Info("Processing notification", "event_type", event.EventType, "user_id", event.UserID)

	return nil
}
