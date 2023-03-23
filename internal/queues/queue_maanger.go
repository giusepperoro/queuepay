package queues

import (
	"fmt"

	ampq "github.com/rabbitmq/amqp091-go"

	"github.com/giusepperoro/queuepay.git/internal/config"
)

const topicNamePrefix = "balance_change_client"

func CreateQueueManager(cfg config.ServiceConfiguration, logger logger) (*QueueManager, error) {
	conn, err := ampq.Dial(cfg.RabbitMQConnectUrl)
	if err != nil {
		_ = conn.Close()
		return nil, err
	}
	return &QueueManager{
		client: conn,
		logger: logger,
	}, nil
}

func (q *QueueManager) makeQueueName(clientId int64) string {
	return fmt.Sprintf("%s_%d", topicNamePrefix, clientId)
}

func (q *QueueManager) makeWorkerName(clientId int64) string {
	return fmt.Sprintf("consume_%s_%d", topicNamePrefix, clientId)
}
