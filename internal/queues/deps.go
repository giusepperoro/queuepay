package queues

import "go.uber.org/zap"

type logger interface {
	Error(msg string, fields ...zap.Field)
}
