package messagebroker

import (
	"github.com/streadway/amqp"
)

func ConnectRabbitMQ() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbit:5672/")
	if err != nil {
		panic(err.Error())
	}

	return conn, nil
}
