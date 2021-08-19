package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io"
	"log"
	"net/http"
)

type TodoItemModel struct {
	ID          int `gorm:"primary_key"`
	Description string
	Completed   bool
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	log.Println("API Health is ok")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	//db, err := gorm.Open("root:root@tcp(127.17.0.2:3306)/todolist?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "root:root@tcp(172.17.0.2:3306)/todolist?parseTime=true")

	if err != nil {
		log.Fatalf("error opening the database %s", err)
	}
	defer db.Close()
	// desc todo_item_models
	db.Debug().DropTableIfExists(&TodoItemModel{})
	db.Debug().AutoMigrate(&TodoItemModel{})
	log.Println("starting todo list api server")
	router := mux.NewRouter()
	router.HandleFunc("/healthz", Healthz).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
