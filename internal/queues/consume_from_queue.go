package queues

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

func (q *QueueManager) ConsumeFromQueue(clientId int64) (<-chan amqp091.Delivery, error) {
	ch, err := q.client.Channel()
	if err != nil {
		q.logger.Error("failed to open channel", zap.Error(err), zap.Int64("client_id", clientId))
		return nil, fmt.Errorf("unable to connect to rabbit: %w", err)
	}
	defer func() {
		_ = ch.Close() // Закрываем канал в случае удачной попытки открытия
	}()

	msgs, err := ch.Consume(
		q.makeQueueName(clientId),
		q.makeWorkerName(clientId),
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		q.logger.Error("failed to  declare a queue", zap.Error(err), zap.Int64("client_id", clientId))
		return msgs, fmt.Errorf("failed to declare a queue: %w", err)
	}

	return msgs, nil
}
