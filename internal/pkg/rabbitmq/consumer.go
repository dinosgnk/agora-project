package rabbitmq

import (
	"fmt"
)

type MessageHandler func([]byte) error

func (c *RabbitMQClient) DeclareQueue(name string) error {
	_, err := c.channel.QueueDeclare(
		name,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	return err
}

func (c *RabbitMQClient) BindQueue(queueName, exchange, routingKey string) error {
	return c.channel.QueueBind(
		queueName,
		routingKey,
		exchange,
		false, // no-wait
		nil,   // arguments
	)
}

func (c *RabbitMQClient) Consume(queueName string, handler MessageHandler) error {
	msgs, err := c.channel.Consume(
		queueName,
		"",    // consumer
		false, // auto-ack (set to false for manual acknowledgment)
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return fmt.Errorf("failed to register consumer: %w", err)
	}

	go func() {
		for msg := range msgs {
			if err := handler(msg.Body); err != nil {
				// Log error but acknowledge to avoid infinite redelivery
				fmt.Printf("Error handling message: %v\n", err)
				msg.Nack(false, false) // Don't requeue
			} else {
				msg.Ack(false)
			}
		}
	}()

	return nil
}
