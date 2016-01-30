package qutils

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const SensorListQueue = "SensorList"

func GetChannel(url string) (*amqp.Connection, *amqp.Channel) {
	log.Print("About to establish connection to queueing system")
	conn, err := amqp.Dial(url)

	if err != nil {
		failOnError(err, "Failed to establish connection to message broker")
	} else {
		log.Printf("Established connection to queue at %s", conn.LocalAddr().String())
	}

	ch, err := conn.Channel()

	if err != nil {
		failOnError(err, "Failed to get channel for connection")
	} else {
		log.Printf("Established connection to queue channel")
	}

	return conn, ch
}

func GetQueue(name string, ch *amqp.Channel) *amqp.Queue {
	q, err := ch.QueueDeclare(
		name,  // name string,
		false, // durable bool,
		false, // autoDelete bool
		false, // exclusive bool,
		false, // noWait bool,
		nil)   // args amqp.Table

	if err != nil {
		failOnError(err, "Failed to declare queue")
	} else {
		log.Printf("Declared [%s] queue successfully", q.Name)
	}

	return &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
