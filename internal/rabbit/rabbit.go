package rabbit

import (
	"github.com/giusepperoro/queuepay.git/internal/config"
	ampq "github.com/rabbitmq/amqp091-go"
	"log"
)

func ConnectRabbit(cfg config.ServiceConfiguration) *QueueRabbit {
	conn, err := ampq.Dial(cfg.RabbitMQConnectUrl)
	if err != nil {
		log.Fatalf("unable to open connect to RabbitMQ server. Error: %s", err)
	}
	defer func() {
		_ = conn.Close()
	}()
	return &QueueRabbit{client: conn}
}
