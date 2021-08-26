package handler

import (
	"encoding/json"
	"errors"
	"exercises/models/api"
	"exercises/models/database"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	storage database.UserStorage
}

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := h.storage.ByID(userId)
	if err != nil {
		log.Println(err)
		if errors.Is(err, database.ErrUserNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	userResponse := api.UserResponseFromDBModel(user)
	if err = json.NewEncoder(w).Encode(userResponse); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h UserHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	var createUserRequest api.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&createUserRequest); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	validate := validator.New()
	if err := validate.Struct(createUserRequest); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := api.UserDBModelFromCreateRequest(createUserRequest)
	if err = h.storage.Add(user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h UserHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updateUserRequest api.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&updateUserRequest); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// checks if request is empty
	emptyRequest := api.UpdateUserRequest{}
	if updateUserRequest == emptyRequest {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// validates json fields
	validate := validator.New()
	if err := validate.Struct(updateUserRequest); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// now gets the id of the user
	user, err := h.storage.ByID(userId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// fills in all updated fields
	if updateUserRequest.Email != nil {
		user.Email = *updateUserRequest.Email
	}

	if updateUserRequest.Names != nil {
		user.Names = *updateUserRequest.Names
	}

	if updateUserRequest.Password != nil {
		hashedPassword, err := api.HashPassword(*updateUserRequest.Password)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user.PasswordHash = hashedPassword
	}

	// saves into database
	if err := h.storage.Update(user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.storage.Delete(userId); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
