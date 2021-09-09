package scheduler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/je117er/tfs-03/lec-11/exercises/models"
	"github.com/je117er/tfs-03/lec-11/exercises/utils"
	"github.com/robfig/cron/v3"
	"github.com/streadway/amqp"
	"log"
	"strings"
)

var (
	USERLENGTH = 10
)

type Producer struct {
	channel      *amqp.Channel
	exchange     string
	exchangeType string
	routingKey   string
}

type Scheduler struct {
	DB       *sql.DB
	c        *cron.Cron
	ctx      context.Context
	producer *Producer
}

func NewScheduler(DB *sql.DB, ch *amqp.Channel, ctx context.Context, exchange string, exchangeType string, routingKey string) *Scheduler {
	return &Scheduler{
		DB:  DB,
		c:   cron.New(cron.WithSeconds()),
		ctx: ctx,
		producer: &Producer{
			channel:      ch,
			exchange:     exchange,
			exchangeType: exchangeType,
			routingKey:   routingKey,
		},
	}
}

func (sched *Scheduler) Start() {
	sched.c.AddFunc("0 * * * * *", sched.scheduleJob)
	log.Println("Starting scheduler...")
	sched.c.Start()
}

func (sched *Scheduler) Stop() {
	log.Println("Stopping sched...")
	sched.c.Stop()
}

func (sched *Scheduler) scheduleJob() {

	// inserts new data to db
	if err := sched.insertData(5); err != nil {
		log.Printf("Error inserting data: %s", err)
		return
	}

	// checks if producer config's correct
	if sched.producer.channel == nil || sched.producer.exchange == "" {
		log.Println("Unknown producer config")
		return
	}

	// scans for new orders
	log.Println("Scanning for new orders...")

	resp, err := sched.getInfoForSending()
	if err != nil {
		return
	}

	// rabbitmq only accepts bytes
	serializedResp, err := utils.Serialize(resp)
	if err != nil {
		return
	}

	log.Println("Scheduling for sending confirmation emails...")
	sched.publish(sched.producer.exchange, sched.producer.routingKey, serializedResp)
}

func (sched *Scheduler) insertData(records int) error {
	sqlStr := "INSERT INTO users (email) VALUES "
	vals := []interface{}{}
	var inserts []string
	for i := 0; i < records; i++ {
		inserts = append(inserts, "(?)")
		vals = append(vals, utils.MakeRandSeq(USERLENGTH)+"@gmail.com")
	}
	sqlStr = sqlStr + strings.Join(inserts, ",")
	_, err := sched.DB.Exec(sqlStr, vals...)
	if err != nil {
		return err
	}
	return nil
}

func (sched *Scheduler) getInfoForSending() ([]*models.InfoResponse, error) {
	resp, err := sched.scanFromDB()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (sched *Scheduler) scanFromDB() ([]*models.InfoResponse, error) {
	rows, err := sched.DB.Query("SELECT id, email FROM users WHERE sent_confirm_email = ?", false)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	defer rows.Close()

	var infoResponses []*models.InfoResponse
	var id uint
	var email string

	for rows.Next() {
		if err := rows.Scan(&id, &email); err != nil {
			continue
		}
		infoResponses = append(infoResponses, &models.InfoResponse{
			ID:    id,
			Email: email,
		})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return infoResponses, nil
}

func (sched *Scheduler) publish(exch, routingKey, body string) error {
	if err := sched.producer.channel.Publish(
		exch,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(body),
			DeliveryMode: amqp.Persistent,
		},
	); err != nil {
		return fmt.Errorf("publishing error: %s", err)
	}
	return nil
}

func (sched *Scheduler) declare() error {
	log.Printf("Binding exchange %v\n", sched.producer.exchange)
	if err := sched.producer.channel.ExchangeDeclare(
		sched.producer.exchange,
		sched.producer.exchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("error binding exchange: %s", err)
	}
	return nil
}

func (sched *Scheduler) Close() error {
	return sched.producer.channel.Close()
}
