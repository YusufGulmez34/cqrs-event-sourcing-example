package main

import (
	"consumer-service/db"
	"encoding/json"
	"log"

	"github.com/couchbase/gocb/v2"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

var collection *gocb.Collection
var items []gocb.BulkOp

func main() {
	collection = db.ConnectCB()

	//RabbitMQ Connect
	conn, err := amqp.Dial("amqp://guest:guest@rabbit:5672/")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	queue, _ := ch.QueueDeclare("useraccount", false, false, false, false, nil)
	msgs, err := ch.Consume(queue.Name, "", true, false, false, false, nil)

	forever := make(chan bool)

	go ListenMessageQueue(msgs)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func ListenMessageQueue(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		err := json.Unmarshal(d.Body, &db.UserAccount)
		if err != nil {
			log.Printf(err.Error())
		}
		_, err = collection.Upsert(uuid.NewString(), db.UserAccount, nil)
		if err != nil {
			log.Println(err)
		}
	}
}
