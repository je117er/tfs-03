package api

import "exercises/models/database"

type cartItemResponse struct {
	ID        int
	CartID    int
	ProductID int
	Price     float64
	Quantity  int8
}

func cartItemResponseFromDBModel(i database.CartItemDBModel) cartItemResponse {
	return cartItemResponse{
		ID:        i.ID,
		CartID:    i.CartID,
		ProductID: i.ProductID,
		Price:     i.Price,
		Quantity:  i.Quantity,
	}
}

type CartItemToOrderItemRequest struct {
	ProductID int
	Price     float64
	Quantity  int8
}
