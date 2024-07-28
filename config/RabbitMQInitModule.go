package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"os"
)

var RabbitMQ *amqp.Connection

func initRabbitMQ() *amqp.Connection {

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s",
		os.Getenv("RABBIT_USER"),
		os.Getenv("RABBIT_PASSWORD"),
		"localhost",
		"5672",
	))

	if err != nil {
		log.Fatalln("RabbitMQ initialization failed\"")
		panic(err)
		return nil
	}

	log.Debug("RabbitMQ initialization succeeded")

	RabbitMQ = conn

	return conn
}
