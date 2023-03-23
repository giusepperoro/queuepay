package handler

import (
	"context"

	"go.uber.org/zap"
)

type logger interface {
	Error(msg string, fields ...zap.Field)
}

type queueProducer interface {
	PutToQueue(ctx context.Context, clientId int64, amount int64) error
}

type workerPool interface {
	RunNewWorker(clientId int64) error
}
