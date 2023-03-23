package queues

import ampq "github.com/rabbitmq/amqp091-go"

type QueueManager struct {
	client *ampq.Connection
	logger logger
}
