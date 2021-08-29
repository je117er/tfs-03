package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"people/application"
	"people/helpers"
	"people/models"
	"strconv"
)

func Home(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			helpers.NotFound(app, w)
			return
		}
		w.Write([]byte("hello"))
	}
}

func CreatePerson(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.Header().Set("Allow", "POST")
			helpers.ClientError(app, w, http.StatusMethodNotAllowed)
			return
		}
		var person models.PersonDBModel
		if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
			helpers.ServerError(app, w, err)
		}
		err := app.PersonStorage.Insert(person)
		if err != nil {
			helpers.ServerError(app, w, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func GetPerson(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			helpers.ClientError(app, w, 400)
		}
		var person *models.PersonDBModel
		person, err = app.PersonStorage.ByID(ID)
		if err != nil {
			helpers.NotFound(app, w)
			return
		}
		if err := json.NewEncoder(w).Encode(person); err != nil {
			helpers.ServerError(app, w, err)
		}
		return
	}
}

func UpdatePerson(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var person models.PersonDBModel
		if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
			helpers.ServerError(app, w, err)
		}
		if err := app.PersonStorage.Update(&person); err != nil {
			helpers.ServerError(app, w, err)
		}
		w.WriteHeader(http.StatusAccepted)
	}
}

func DeletePerson(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.ParseUint(vars["id"], 10, 32)
		if err != nil {
			helpers.ClientError(app, w, 400)
		}
		if err := app.PersonStorage.Delete(ID); err != nil {
			helpers.ServerError(app, w, err)
		}
		w.WriteHeader(http.StatusAccepted)
	}
}

func GetPeople(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var people []models.PersonDBModel
		people, err := app.PersonStorage.All()
		if err != nil {
			helpers.ServerError(app, w, err)
		}

		if err := json.NewEncoder(w).Encode(people); err != nil {
			helpers.ServerError(app, w, err)
		}
		return
	}
}
