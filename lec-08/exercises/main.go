package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/je117er/tfs-03/lec-08/exercises/config"
	"github.com/olivere/elastic/v7"
	"log"
	"strconv"
	"time"
)

var (
	searchTerm = "this book has changed my life!"
	ctx        = context.Background()
)

func mysqlElapsed() time.Duration {
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Println("error parsing db port")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USERNAME"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	matchTerm := "%" + searchTerm + "%"
	start := time.Now()
	rows, err := db.QueryContext(ctx, "SELECT * FROM reviews WHERE body LIKE ?", matchTerm)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return time.Since(start)
}

func esElapsed() float64 {
	// starts a new es client
	client, err := elastic.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	// pings es server to get version number
	info, code, err := client.Ping("http://localhost:9200").Do(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Es returned with code %d and version %s\n", code, info.Version.Number)

	// checks if an index exists
	_, err = client.IndexExists("reviews").Do(ctx)
	if err != nil {
		log.Fatalf("Index %s doesn't exist", err)
	}
	fmt.Println("Index exists!")

	// searches for a term
	termQuery := elastic.NewMatchQuery("body", searchTerm)
	searchResult, err := client.Search().
		Index("reviews").
		Query(termQuery).
		Pretty(true).Do(ctx)
	if err != nil {
		log.Printf("error %s occured while searching", err)
	}

	// by default the number of hits is limited to 10,000
	// reference: https://www.elastic.co/guide/en/elasticsearch/reference/current/paginate-search-results.html
	log.Printf("Found %d results", searchResult.TotalHits())

	return float64(searchResult.TookInMillis) / 1000
}

func main() {

	mysqlElapsed := mysqlElapsed()
	esElapsed := esElapsed()
	log.Printf("Query by MySQL took %s\n", mysqlElapsed)
	log.Printf("Query by Elasticsearch took %fs\n", esElapsed)
}
