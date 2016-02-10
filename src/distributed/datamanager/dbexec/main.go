package main

import (
	"bytes"
	"distributed/datamanager"
	"distributed/dto"
	"distributed/qutils"
	"encoding/gob"
	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

const url = "amqp://guest:guest@localhost:5672"

func main() {
	conn, ch := qutils.GetChannel(url)
	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(
		qutils.PersistReadingsQueue, //queue string,
		"",    //consumer string,
		false, //autoAck bool,
		true,  //exclusive bool,
		false, //noLocal bool,
		false, //noWait bool,
		nil)   //args amqp.Table)

	if err != nil {
		log.Warn("Failed to get access to messages")
	}
	for msg := range msgs {
		buf := bytes.NewReader(msg.Body)
		dec := gob.NewDecoder(buf)
		sd := &dto.SensorMessage{}
		dec.Decode(sd)

		err := datamanager.SaveReading(sd)

		if err != nil {
			log.Error("Failed to save reading from sensor", sd.Name,". Error: ", err.Error())
		} else {
			msg.Ack(false)
		}
	}
}
