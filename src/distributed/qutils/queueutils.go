package qutils

import (
	"fmt"
	"github.com/streadway/amqp"
	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

const SensorListQueue = "SensorList"

const SensorDiscoveryExchange = "SensoryDiscovery"
const PersistReadingsQueue = "PersistReading"

func GetChannel(url string) (*amqp.Connection, *amqp.Channel) {
	log.Info("About to establish connection to queueing system")
	conn, err := amqp.Dial(url)

	if err != nil {
		failOnError(err, "Failed to establish connection to message broker")
	} else {
		log.Info("Established connection to queue at ", conn.LocalAddr().String())
	}

	ch, err := conn.Channel()

	if err != nil {
		failOnError(err, "Failed to get channel for connection")
	} else {
		log.Info("Established connection to queue channel")
	}

	return conn, ch
}

func GetQueue(name string, ch *amqp.Channel, autoDelete bool) *amqp.Queue {
	q, err := ch.QueueDeclare(
		name,  // name string,
		false, // durable bool,
		autoDelete, // autoDelete bool
		false, // exclusive bool,
		false, // noWait bool,
		nil)   // args amqp.Table

	if err != nil {
		failOnError(err, "Failed to declare queue")
	} else {
		log.Info("Declared", q.Name,"queue successfully")
	}

	return &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Error(msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
