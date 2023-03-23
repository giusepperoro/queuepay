package queues

import (
	"fmt"
	"github.com/giusepperoro/queuepay.git/internal/config"
	ampq "github.com/rabbitmq/amqp091-go"
)

const topicNamePrefix = "balance_change_client"
const topicTestForOne = "balance_change_client_1"

func CreateQueueManager(cfg config.ServiceConfiguration, logger logger) (*QueueManager, error) {
	conn, err := ampq.Dial(cfg.RabbitMQConnectUrl)
	if err != nil {
		_ = conn.Close()
		return nil, err
	}
	go createConsumer(conn)
	return &QueueManager{
		client: conn,
		logger: logger,
	}, nil
}

func createConsumer(conn *ampq.Connection) {
	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("failed to declare a queue: %e", err)
		return
	}
	msgs, err := ch.Consume(
		topicTestForOne, // queue
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		fmt.Printf("Failed to register a consumer: %e", err)
		return
	}

	go func() {
		for d := range msgs {
			fmt.Printf("Received a message: %s", d.Body)
		}
	}()

}
