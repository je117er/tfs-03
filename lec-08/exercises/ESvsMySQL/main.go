package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/je117er/tfs-03/lec-08/exercises/config"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	ctx = context.Background()
)

func mysqlElapsed(searchTerm string) time.Duration {
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
	elapsed := time.Since(start)
	defer rows.Close()

	return elapsed
}

func esElapsed(searchTerm string) time.Duration {
	// starts a new es client
	client, err := elastic.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// searches for a term
	termQuery := elastic.NewMatchPhraseQuery("body", searchTerm)
	start := time.Now()
	_, err = client.Search().
		Index("reviews").
		Query(termQuery).
		Do(ctx)
	if err != nil {
		log.Printf("error %s occured while searching", err)
	}
	elapsed := time.Since(start)
	//log.Println(time.Since(start))
	// by default the number of hits is limited to 10,000
	// reference: https://www.elastic.co/guide/en/elasticsearch/reference/current/paginate-search-results.html
	//log.Printf("Found %d results", searchResult.TotalHits())

	return elapsed
}

func printHelper(engineChoice int, elapsed string, searchTerm string) {
	var engine string
	switch engineChoice {
	case 1:
		engine = "MySQL"
	case 2:
		engine = "Elasticsearch"
	}
	log.Printf("Query for %q by %s took %s", searchTerm, engine, elapsed)
}

func main() {
	searchTerm := strings.Join(os.Args[1:], " ")
	for i := 0; i < 10; i++ {
		printHelper(1, mysqlElapsed(searchTerm).String(), searchTerm)
	}
	for i := 0; i < 10; i++ {
		printHelper(2, esElapsed(searchTerm).String(), searchTerm)
	}
}
