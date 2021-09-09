package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/je117er/tfs-03/lec-11/exercises/config"
	"github.com/je117er/tfs-03/lec-11/exercises/db"
	"github.com/je117er/tfs-03/lec-11/exercises/rmqClient"
	"github.com/je117er/tfs-03/lec-11/exercises/scheduler"
	"github.com/je117er/tfs-03/lec-11/exercises/worker"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	DB, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

	// initializes rabbitmq client
	rmqSource := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.Config("RMQ_USERNAME"), config.Config("RMQ_PASSWORD"), config.Config("RMQ_HOST"), config.Config("RMQ_PORT"))
	client := rmqClient.NewRMQClient(rmqSource)

	// two channels for two actors
	sCh, err := client.GetChannel()
	if err != nil {
		log.Fatalf("Unable to get channel: %s", err)
	}
	wCh, err := client.GetChannel()
	if err != nil {
		log.Fatalf("Unable to get channel: %s", err)
	}

	// declares the rest of the parameters
	exch := "order"
	exchType := ""
	queue := "order_processor"
	routingKey := ""
	var wg *sync.WaitGroup
	ctx, cancelFunc := context.WithCancel(context.Background())

	sched := scheduler.NewScheduler(DB, sCh, ctx, exch, exchType, routingKey)
	wkr := worker.NewWorker(wg, ctx, DB, wCh, queue, exch, exchType, routingKey)

	// gracefully shuts down
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		sig := <-c
		log.Printf("Got %s signal. Starting shutting down...", sig)
		sched.Close()
		wkr.Close()
		cancelFunc()
	}()

	wg.Add(1)
	go wkr.Start()

	sched.Start()

	wg.Wait()
}
