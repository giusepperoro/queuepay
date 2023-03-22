package rabbit

import ampq "github.com/rabbitmq/amqp091-go"

type QueueRabbit struct {
	Client *ampq.Connection
}
