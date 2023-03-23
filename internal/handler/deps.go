package handler

import "go.uber.org/zap"

type logger interface {
	Error(msg string, fields ...zap.Field)
}

type queueProducer interface {
	PutToQueue(clientId, amount int64) error
}
