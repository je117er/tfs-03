package handler

import (
	"encoding/json"
	"exercises/config"
	"exercises/models/api"
	"exercises/models/database"
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"log"
	"net/http"
)

type PaymentHandler struct {
	storage database.PaymentStorage
}

func (h PaymentHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var createPaymentRequest api.CreatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&createPaymentRequest); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	stripe.Key = config.Config("STRIPE_API_KEY")
	orderIDString := fmt.Sprintf("OrderID %d", createPaymentRequest.OrderID)

	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(int64(createPaymentRequest.Amount)),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Description: stripe.String(orderIDString),
		Source:      &stripe.SourceParams{Token: stripe.String("tok_visa")},
	}
	if _, err := charge.New(params); err != nil {
		log.Println(err)
		http.Error(w, "Payment request failed", http.StatusInternalServerError)
		return
	}
	payment := api.PaymentDBModelFromCreateRequest(createPaymentRequest)
	if err := h.storage.Add(payment); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
