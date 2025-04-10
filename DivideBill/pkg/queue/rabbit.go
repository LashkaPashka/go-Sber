package queue

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/lashkapashka/divideBill/internal/model"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	ch *amqp.Channel
}

func failOnError(err error, msg string) {
	if err != nil {
	  log.Panicf("%s: %s", msg, err)
	}
  }

func New() *RabbitMQ{
	var rabbit RabbitMQ

	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	rabbit.ch = ch
	
	return &rabbit
}

func (r *RabbitMQ) Producer(msg model.Response) {
	q, err := r.ch.QueueDeclare(
		"topic-divide",
		false,
		false,
		false,   
		false,   
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	msgValue, err := json.Marshal(msg)
	failOnError(err, "couldn't convert the message")

	err = r.ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: msgValue,
		})
	failOnError(err, "Failed to publish a message")
	log.Println("сообщение отправлено в RabbitMQ")
}

func (r *RabbitMQ) Consumer() {
	q, err := r.ch.QueueDeclare(
		"topic-divide",
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	msgs, err := r.ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	
	go func() {
		for d := range msgs {
		  log.Printf("Received a message: %s", d.Body)
		}
	  }()
	
	log.Println("Waiting for messages.")
	
	select{}
}