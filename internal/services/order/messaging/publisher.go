package messaging

import (
	"fmt"
	"time"

	"github.com/dinosgnk/agora-project/internal/pkg/rabbitmq"
	"github.com/google/uuid"
)

const (
	OrderExchange = "orders"
)

type Publisher struct {
	client *rabbitmq.RabbitMQClient
}

func NewPublisher(client *rabbitmq.RabbitMQClient) (*Publisher, error) {
	if err := client.DeclareExchange(OrderExchange, "topic"); err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	return &Publisher{
		client: client,
	}, nil
}

func (p *Publisher) PublishOrderCreated(event *OrderCreatedEvent) error {
	event.EventID = uuid.New().String()
	event.Timestamp = time.Now()
	routingKey := "order.created"

	return p.client.PublishMessage(OrderExchange, routingKey, event)
}

func (p *Publisher) PublishOrderStatusUpdated(event *OrderStatusUpdatedEvent) error {
	event.EventID = uuid.New().String()
	event.Timestamp = time.Now()
	routingKey := "order.status.updated"

	return p.client.PublishMessage(OrderExchange, routingKey, event)
}

func (p *Publisher) PublishOrderConfirmed(event *OrderConfirmedEvent) error {
	event.EventID = uuid.New().String()
	event.Timestamp = time.Now()
	routingKey := "order.confirmed"

	return p.client.PublishMessage(OrderExchange, routingKey, event)
}

func (p *Publisher) PublishOrderProcessing(event *OrderProcessingEvent) error {
	event.EventID = uuid.New().String()
	event.Timestamp = time.Now()
	routingKey := "order.processing"

	return p.client.PublishMessage(OrderExchange, routingKey, event)
}

func (p *Publisher) PublishOrderShipped(event *OrderShippedEvent) error {
	event.EventID = uuid.New().String()
	event.Timestamp = time.Now()
	routingKey := "order.shipped"

	return p.client.PublishMessage(OrderExchange, routingKey, event)
}

func (p *Publisher) PublishOrderDelivered(event *OrderDeliveredEvent) error {
	event.EventID = uuid.New().String()
	event.Timestamp = time.Now()
	routingKey := "order.delivered"

	return p.client.PublishMessage(OrderExchange, routingKey, event)
}

func (p *Publisher) PublishOrderCancelled(event *OrderCancelledEvent) error {
	event.EventID = uuid.New().String()
	event.Timestamp = time.Now()
	routingKey := "order.cancelled"

	return p.client.PublishMessage(OrderExchange, routingKey, event)
}
