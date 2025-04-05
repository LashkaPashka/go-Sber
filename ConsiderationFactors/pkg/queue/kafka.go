package queue

import (
	"context"
	"errors"
	"log"

	"github.com/lashkapshka/go-Sber/internal/service"
	"github.com/segmentio/kafka-go"
)

type Client struct {
	Reader *kafka.Reader
	Writer *kafka.Writer
	service *service.FactorsService
}

func New(brokers []string, topic string, groupId string) (*Client, error) {
	if len(brokers) == 0 || brokers[0] == "" || topic == "" || groupId == ""{
		return nil, errors.New("не указаны параметры подключения к Kafka")
	}
	
	c := Client{}

	c.Reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupId,
		MinBytes: 10e1,
		MaxBytes: 10e6,
	})

	c.Writer = &kafka.Writer{
		Addr:                   kafka.TCP(brokers[0]),
		Topic:                  topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}


	return &c, nil
}

func (c *Client) Producer(msgVal string) {
	msg := kafka.Message{
		Value: []byte(msgVal),
	}

	err := c.Writer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Println(err)
	}
}

func (c *Client) Consumer() {
	for {
		msg, err := c.Reader.FetchMessage(context.Background())
		if err != nil {
			log.Println(err)
		}

		log.Println(string(msg.Value))
		c.service.DivideBill(string(msg.Value))

		err = c.Reader.CommitMessages(context.Background(), msg)
		if err != nil {
			log.Println(err)
		}
	}
}