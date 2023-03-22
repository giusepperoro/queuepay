package rabbit

import (
	"context"
	"github.com/giusepperoro/queuepay.git/internal/config"
	ampq "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func ConnectRabbit(cfg config.ServiceConfiguration) (*QueueRabbit, error) {
	conn, err := ampq.Dial(cfg.RabbitMQConnectUrl)
	if err != nil {
		_ = conn.Close()
		return nil, err
	}
	return &QueueRabbit{Client: conn}, nil
}

func ClientQueue(conn *ampq.Connection, clientId, amount int64) {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open channel. Error: %s", err)
	}

	defer func() {
		_ = ch.Close() // Закрываем канал в случае удачной попытки открытия
	}()
	q, err := ch.QueueDeclare(
		string(clientId), // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		log.Fatalf("failed to declare a queue. Error: %s", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := string(amount)
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		ampq.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("failed to publish a message. Error: %s", err)
	}

	log.Printf(" [x] Sent %s\n", body)
}
