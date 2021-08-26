package handler

import (
	"encoding/json"
	"exercises/internal/models/api"
	"exercises/internal/models/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	storage database.OrderStorage
}

func (h OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	var ordersResponse []api.OrderResponse
	orders, err := h.storage.All()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, e := range orders {
		ordersResponse = append(ordersResponse, api.OrderResponseFromDBModel(e))
	}

	if err := json.NewEncoder(w).Encode(ordersResponse); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	order, err := h.storage.ByID(orderID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	orderResponse := api.OrderResponseFromDBModel(order)
	if err := json.NewEncoder(w).Encode(orderResponse); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
