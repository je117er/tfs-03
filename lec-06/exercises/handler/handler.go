package handler

import (
	"encoding/json"
	"errors"
	"exercises/models"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	storage models.UserStorage
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
		if errors.Is(err, models.ErrUserNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	userResponse := models.UserResponseFromDBModel(user)
	if err = json.NewEncoder(w).Encode(userResponse); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h UserHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	var createUserRequest models.CreateUserRequest
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

	user, err := models.UserDBModelFromCreateRequest(createUserRequest)
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

	var updateUserRequest models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&updateUserRequest); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// checks if request is empty
	emptyRequest := models.UpdateUserRequest{}
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
		hashedPassword, err := models.HashPassword(*updateUserRequest.Password)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user.PasswordHash = hashedPassword
	}

	// saves into db
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

type ProductHandler struct {
	storage models.ProductStorage
}

func (h ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.storage.All()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var productsResponse []models.ProductResponse
	for _, p := range products {
		productsResponse = append(productsResponse, models.ProductResponseFromDBModel(p))
	}

	if err := json.NewEncoder(w).Encode(productsResponse); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.storage.ByID(productID)
	if err != nil {
		log.Println(err)
		if errors.Is(err, models.ErrProductNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	productResponse := models.ProductResponseFromDBModel(product)
	if err := json.NewEncoder(w).Encode(productResponse); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h ProductHandler) PostProduct(w http.ResponseWriter, r *http.Request) {
	var createProductRequest models.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&createProductRequest); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	validate := validator.New()
	if err := validate.Struct(createProductRequest); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product := models.ProductDBModelFromCreateRequest(createProductRequest)
	if err := h.storage.Add(product); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h ProductHandler) PatchProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var updateProductRequest models.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&updateProductRequest); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	emptyRequest := models.UpdateProductRequest{}
	if emptyRequest == updateProductRequest {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	validate := validator.New()
	if err := validate.Struct(updateProductRequest); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.storage.ByID(productID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if updateProductRequest.Title != nil {
		product.Title = *updateProductRequest.Title
	}
	if updateProductRequest.Price != nil {
		product.Price = *updateProductRequest.Price
	}

	if updateProductRequest.Description != nil {
		product.Description = *updateProductRequest.Description
	}

	if updateProductRequest.Quantity != nil {
		product.Quantity = *updateProductRequest.Quantity
	}

	if updateProductRequest.Status != nil {
		product.Status = *updateProductRequest.Status
	}

	if err := h.storage.Update(product); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.storage.Delete(productID); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
