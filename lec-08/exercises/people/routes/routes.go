package routes

import (
	"github.com/gorilla/mux"
	"people/application"
	"people/handlers"
)

func Routes(app *application.Application) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Home(app)).Methods("GET")
	r.HandleFunc("/person", handlers.CreatePerson(app)).Methods("POST")
	r.HandleFunc("/person/{id:[0-9]+}", handlers.GetPerson(app)).Methods("GET")
	r.HandleFunc("/person/{id:\\d+}", handlers.UpdatePerson(app)).Methods("PUT")
	r.HandleFunc("/person/{id:\\d+}", handlers.DeletePerson(app)).Methods("DELETE")
	r.HandleFunc("/people", handlers.GetPeople(app)).Methods("GET")
	return r
}
