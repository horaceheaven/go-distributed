package coordinator

import "github.com/streadway/amqp"

const url = "amqp:guest:guest@localhost:5672"

type QueueListener struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}
