package api

import (
	"exercises/internal/models/database"
	"time"
)

type CreatePaymentRequest struct {
	OrderID      int
	Amount       float64
	ReceiptEmail string
}

func PaymentDBModelFromCreateRequest(r CreatePaymentRequest) database.PaymentDBModel {
	return database.PaymentDBModel{
		OrderID:      r.OrderID,
		Amount:       r.Amount,
		ReceiptEmail: r.ReceiptEmail,
		PaymentDate:  time.Now(),
	}
}
