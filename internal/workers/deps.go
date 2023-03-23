package workers

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type logger interface {
	Error(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
}

type queueManager interface {
	ConsumeFromQueue(clientId int64) (<-chan amqp091.Delivery, error)
}

type database interface {
	ChangeBalance(ctx context.Context, clientId int64, amount int64) error
}
