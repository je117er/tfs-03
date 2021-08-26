package api

import "exercises/models/database"

type CreateCartRequest struct {
	UserID int
	Status string
	Total  int
}

type ProductToCartRequest struct {
	CartID   int
	ID       int
	Price    float64
	Quantity int8
}

func ProductToCartItemDBModelRequest(p ProductToCartRequest) database.CartItemDBModel {
	return database.CartItemDBModel{
		CartID:    p.CartID,
		Price:     p.Price,
		Quantity:  p.Quantity,
		ProductID: p.ID,
	}
}

type UpdateCartRequest struct {
	Status   *string
	Products *[]ProductToCartRequest
}

type CartResponse struct {
	ID        int
	UserID    int
	CartItems []cartItemResponse
	Status    string
	Total     float64
}

func CartResponseFromDBModel(c database.CartDBModel) CartResponse {
	var cartItems []cartItemResponse
	for _, e := range c.CartItems {
		cartItems = append(cartItems, cartItemResponseFromDBModel(e))
	}

	return CartResponse{
		ID:        c.ID,
		UserID:    c.UserID,
		Status:    c.Status,
		CartItems: cartItems,
		Total:     c.Total,
	}
}

func CartDBModelFromCreateRequest(r CreateCartRequest) database.CartDBModel {
	return database.CartDBModel{
		UserID: r.UserID,
		Total:  0,
		Status: r.Status,
	}
}

type CreateCheckoutRequest struct {
	CartID    int
	UserID    int
	CartItems []CartItemToOrderItemRequest
	Total     float64
}

func OrderDBModelFromCreateCheckOutRequest(r CreateCheckoutRequest) database.OrderDBModel {
	var orderItems []database.OrderItemDBModel
	for _, e := range r.CartItems {
		orderItems = append(orderItems, CartToOrderItemDBModelRequest(e))
	}
	return database.OrderDBModel{
		UserID:     r.UserID,
		Total:      r.Total,
		OrderItems: orderItems,
		Status:     "Unpaid",
	}
}

func CartToOrderItemDBModelRequest(r CartItemToOrderItemRequest) database.OrderItemDBModel {
	return database.OrderItemDBModel{
		ProductID: r.ProductID,
		Price:     r.Price,
		Quantity:  r.Quantity,
	}
}
