package main

import (
	"log"
	"net/http"
	"os"
	"people/application"
	"people/helpers"
	"people/models"
	"people/routes"
)

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// Writes error messages to stderr
	// log.Lshortfile will forward file name
	// and line number to stderr
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := helpers.OpenDB(helpers.GetDSN())
	if err != nil {
		errorLog.Fatal(err)
	}

	if err := helpers.CreateTable(db); err != nil {
		errorLog.Fatal("failed to create tables")
	}
	infoLog.Println("successfully created tables")

	app := &application.Application{
		ErrorLog:      errorLog,
		InfoLog:       infoLog,
		PersonStorage: &models.PersonStorage{DB: db},
	}

	// Initializes a new http.Server struct
	// using custom ErrorLog
	srv := &http.Server{
		Addr:     ":8000",
		ErrorLog: errorLog,
		Handler:  routes.Routes(app),
	}
	infoLog.Printf("Starting server on :8000")
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
