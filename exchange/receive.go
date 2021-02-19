package exchange

import (

	"github.com/streadway/amqp"
	"time"
)


func Receive() ([]byte, error){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	checkerror(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	checkerror(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"eastern",
		false,
		false,
		false,
		false,
		nil,
	)
	checkerror(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	checkerror(err, "Failed to register a consumer")
	var msg []byte
	go func() {
		for d := range msgs {
			msg = d.Body
		}
	}()
	time.Sleep(1 * time.Second)
	return msg ,nil

}