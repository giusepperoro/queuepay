package workers

import (
	"context"
	"strconv"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

func (w *worker) Consume(msgs <-chan amqp091.Delivery, clientId int64) {
	for msg := range msgs {
		amount, err := strconv.ParseInt(string(msg.Body), 10, 64)
		if err != nil {
			w.logger.Error("unable to get amount", zap.Error(err), zap.Int64("clientId", clientId))
		}

		err = w.changeDataInDatabase(clientId, amount)
		if err != nil {
			w.logger.Error("unable to change balance", zap.Error(err), zap.Int64("clientId", clientId))
		} else {
			err = msg.Ack(true)
			if err != nil {
				w.logger.Error("unable to consume", zap.Error(err), zap.Int64("clientId", clientId))
			}
		}
	}
	w.logger.Info("worker finished reading from queue", zap.Int64("clientId", clientId))
}

func (w *worker) changeDataInDatabase(clientId, amount int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return w.database.ChangeBalance(ctx, clientId, amount)
}
