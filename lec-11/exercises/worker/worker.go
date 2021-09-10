package worker

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/je117er/tfs-03/lec-11/exercises/models"
	"github.com/je117er/tfs-03/lec-11/exercises/utils"
	"github.com/streadway/amqp"
	"log"
	"strings"
	"sync"
)

type Consumer struct {
	chann      *amqp.Channel
	queue      string
	exch       string
	exchType   string
	bindingKey string
}

type Worker struct {
	wg       *sync.WaitGroup
	ctx      context.Context
	DB       *sql.DB
	consumer *Consumer
}

func NewWorker(wg *sync.WaitGroup, ctx context.Context, DB *sql.DB, chann *amqp.Channel, queue string, exch string, exchType string, bindingKey string) *Worker {
	return &Worker{
		wg:  wg,
		ctx: ctx,
		DB:  DB,
		consumer: &Consumer{
			chann:      chann,
			queue:      queue,
			exch:       exch,
			exchType:   exchType,
			bindingKey: bindingKey,
		},
	}
}

func (w *Worker) Start() {
	if w.consumer.chann == nil || w.consumer.queue == "" {
		log.Println("Incorrect consumer config")
	}

	if err := w.declare(); err != nil {
		log.Println(err)
		return
	}

	log.Println("Queue is now bound to exchange. Starting to consume data...")
	delivery, err := w.consumer.chann.Consume(
		w.consumer.queue,
		// consumer identifier, "" means it'll be autogenerated
		// references: https://www.rabbitmq.com/amqp-0-9-1-reference.html#basic.consume.consumer-tag
		"",
		false,
		false,
		// can send to consumers on the same connection
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("error consuming queue: %s", err)
		return
	}

	for {
		select {
		case msg := <-delivery:
			log.Println("Message received")
			// only acknowledges this delivery
			msg.Ack(false)

			// body containing a list of info
			body, err := utils.Deserialize(msg.Body)
			if err != nil {
				log.Println(err)
			}
			if err := w.updateDB(body); err != nil {
				log.Println(err)
			}
		case <-w.ctx.Done():
			log.Println("Worker exiting...")
			w.wg.Done()
			return
		}
	}
}

func (w *Worker) updateDB(resp []*models.InfoResponse) error {
	sqlStr := "INSERT INTO users (email) VALUES "
	vals := []interface{}{}
	var inserts []string
	for _, info := range resp {
		inserts = append(inserts, "(?) (?) (?) ")
		vals = append(vals, info.ID, info.Email, true)
	}
	sqlStr = sqlStr + strings.Join(inserts, ",") + "\n"
	sqlStr = sqlStr + `
		ON DUPLICATE KEY UPDATE 
		email = VALUES(email),
		sent_confirm_email = VALUES(sent_confirm_email)
`
	_, err := w.DB.Exec(sqlStr, vals...)
	if err != nil {
		return err
	}

	return nil
}

func (w *Worker) declare() error {
	log.Printf("Binding exchange %v\n", w.consumer.exch)
	if err := w.consumer.chann.ExchangeDeclare(
		w.consumer.exch,
		w.consumer.exchType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("error declaring exchange: %s", err)
	}

	log.Printf("Declaring queue %v\n", w.consumer.queue)
	queue, err := w.consumer.chann.QueueDeclare(
		// make sure a queue with same name doesn't exist
		// queue name is automatically generated
		// when none is specified
		w.consumer.queue,
		// whether queue survives broker's restart
		true,
		// remove when last consumer is deleted nor not?
		false,
		// used by one connection or not?
		false,
		false,
		// for broker plugins and other features
		nil,
	)
	if err != nil {
		return fmt.Errorf("error declaring queue: %s", err)
	}

	log.Printf("Binding exchange %v to queue %v\n", w.consumer.exch, w.consumer.queue)
	if err := w.consumer.chann.QueueBind(
		queue.Name,
		w.consumer.bindingKey,
		w.consumer.exch,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("error binding queue: %s", err)
	}

	return nil
}

func (w *Worker) Close() error {
	return w.consumer.chann.Close()
}
