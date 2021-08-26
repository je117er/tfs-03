package api

import (
	"exercises/internal/models/database"
)

type OrderItemResponse struct {
	ID        int
	ProductID int
	Price     float64
	Quantity  int8
}

func OrderItemResponseFromDBModel(o database.OrderItemDBModel) OrderItemResponse {
	return OrderItemResponse{
		ID:        o.ID,
		ProductID: o.ProductID,
		Price:     o.Price,
		Quantity:  o.Quantity,
	}
}
