package handlers

import (
	"encoding/json"
	"errors"
	"exercises/internal/models/api"
	database2 "exercises/internal/models/database"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	storage database2.ProductStorage
}

func (h ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.storage.All()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var productsResponse []api.ProductResponse
	for _, p := range products {
		productsResponse = append(productsResponse, api.ProductResponseFromDBModel(p))
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
		if errors.Is(err, database2.ErrProductNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	productResponse := api.ProductResponseFromDBModel(product)
	if err := json.NewEncoder(w).Encode(productResponse); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h ProductHandler) PostProduct(w http.ResponseWriter, r *http.Request) {
	var createProductRequest api.CreateProductRequest
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

	product := api.ProductDBModelFromCreateRequest(createProductRequest)
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
	var updateProductRequest api.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&updateProductRequest); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	emptyRequest := api.UpdateProductRequest{}
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
