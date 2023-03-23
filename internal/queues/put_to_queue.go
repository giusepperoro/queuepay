package queues

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	ampq "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

func (q *QueueManager) PutToQueue(clientId, amount int64) error {
	ch, err := q.client.Channel()
	if err != nil {
		q.logger.Error("failed to open channel", zap.Error(err), zap.Int64("client_id", clientId))
		return fmt.Errorf("unable to connect to rabbit: %w", err)
	}
	defer func() {
		_ = ch.Close() // Закрываем канал в случае удачной попытки открытия
	}()

	queue, err := ch.QueueDeclare(
		fmt.Sprintf("%s_%d", topicNamePrefix, clientId), // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		q.logger.Error("failed to  declare a queue", zap.Error(err), zap.Int64("client_id", clientId))
		return fmt.Errorf("failed to declare a queue: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := []byte(strconv.FormatInt(amount, 10))
	err = ch.PublishWithContext(ctx,
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		ampq.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		q.logger.Error("failed to publish", zap.Error(err), zap.Int64("client_id", clientId))
		return fmt.Errorf("unable to publish: %w", err)
	}

	log.Printf(" [x] Sent %s\n", body)
	return nil
}
