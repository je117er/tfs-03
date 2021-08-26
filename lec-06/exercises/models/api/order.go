package api

import "exercises/models/database"

type OrderResponse struct {
	ID         int
	UserID     int
	OrderItems []OrderItemResponse
	Total      float64
}

func OrderResponseFromDBModel(o database.OrderDBModel) OrderResponse {
	var orderItems []OrderItemResponse
	for _, e := range o.OrderItems {
		orderItems = append(orderItems, OrderItemResponseFromDBModel(e))
	}
	return OrderResponse{
		ID:         o.ID,
		UserID:     o.UserID,
		Total:      o.Total,
		OrderItems: orderItems,
	}
}
