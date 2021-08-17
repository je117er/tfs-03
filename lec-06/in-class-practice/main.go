package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	ID   int
	Name string
}

var students []Student

const (
	user = "user"
	pass = "pass"
)

const JSONContentType = "application/json"

func contentTypeCheckMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")

		if reqContentType != JSONContentType {
			fmt.Fprintf(w, "only allows application/json")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if username, password, ok := r.BasicAuth(); ok {
			suppliedUsername := username
			suppliedPassword := password
			if suppliedUsername == user && suppliedPassword == pass {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "You don't have access to this page", http.StatusUnauthorized)
	})
}

func main() {

	students = append(students, Student{1, "j"})
	students = append(students, Student{2, "jo"})
	r := mux.NewRouter()
	r.HandleFunc("/students", getStudents).Methods("GET")
	r.HandleFunc("/students", createStudent).Methods("POST")
	r.Use(Auth)
	r.Use(contentTypeCheckMiddleWare)

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	students = append(students, student)
	json.NewEncoder(w).Encode(&students)
}
