package rmqClient

import (
	"github.com/streadway/amqp"
	"log"
)

type RMQClient struct {
	URL  string
	conn *amqp.Connection
}

func NewRMQClient(URL string) *RMQClient {
	conn, err := amqp.Dial(URL)
	if err != nil {
		log.Printf("Unable to start rabbitmq: %s", err)
	}
	return &RMQClient{
		URL:  URL,
		conn: conn,
	}
}

func (client *RMQClient) GetChannel() (*amqp.Channel, error) {
	return client.conn.Channel()
}

func (client *RMQClient) Close() {
	client.conn.Close()
}
