package models

import (
	_ "github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	Username string
	Password string
	Email    string `json:"email" validate:"required,email"`
	Names    string
}

type UpdateUserRequest struct {
	Email    *string
	Names    *string
	Password *string
}

type UserResponse struct {
	ID    int
	Email string
	Names string
}

func UserResponseFromDBModel(u UserDBModel) UserResponse {
	return UserResponse{
		ID:    u.ID,
		Email: u.Email,
		Names: u.Names,
	}
}

func UserDBModelFromCreateRequest(r CreateUserRequest) (UserDBModel, error) {
	hashedPassword, err := HashPassword(r.Password)
	if err != nil {
		return UserDBModel{}, err
	}
	return UserDBModel{
		Username:     r.Username,
		Email:        r.Email,
		PasswordHash: hashedPassword,
		Names:        r.Names,
	}, nil
}

type CreateProductRequest struct {
	Title       string
	Description string
	Price       float64 `validate:"numeric"`
	Quantity    int8    `validate:"numeric"`
}

func ProductDBModelFromCreateRequest(r CreateProductRequest) ProductDBModel {
	return ProductDBModel{
		Title:       r.Title,
		Description: r.Description,
		Price:       r.Price,
		Quantity:    r.Quantity,
		Status:      "In Stock",
	}
}

type UpdateProductRequest struct {
	Title       *string
	Description *string
	Price       *float64 `validate:"numeric"`
	Quantity    *int8    `validate:"numeric"`
	Status      *string
}

type ProductResponse struct {
	ID          int
	Title       string
	Description string
	Price       float64
	Quantity    int8
	Status      string
}

func ProductResponseFromDBModel(p ProductDBModel) ProductResponse {
	return ProductResponse{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    p.Quantity,
		Status:      p.Status,
	}
}

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

func ProductToCartItemDBModelRequest(p ProductToCartRequest) CartItemDBModel {
	return CartItemDBModel{
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

type cartItemResponse struct {
	ID        int
	CartID    int
	ProductID int
	Price     float64
	Quantity  int8
}

func cartItemResponseFromDBModel(i CartItemDBModel) cartItemResponse {
	return cartItemResponse{
		ID:        i.ID,
		CartID:    i.CartID,
		ProductID: i.ProductID,
		Price:     i.Price,
		Quantity:  i.Quantity,
	}
}

func CartResponseFromDBModel(c CartDBModel) CartResponse {
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

func CartDBModelFromCreateRequest(r CreateCartRequest) CartDBModel {
	return CartDBModel{
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

type CartItemToOrderItemRequest struct {
	ProductID int
	Price     float64
	Quantity  int8
}

func OrderDBModelFromCreateCheckOutRequest(r CreateCheckoutRequest) OrderDBModel {
	var orderItems []OrderItemDBModel
	for _, e := range r.CartItems {
		orderItems = append(orderItems, CartToOrderItemDBModelRequest(e))
	}
	return OrderDBModel{
		UserID:     r.UserID,
		Total:      r.Total,
		OrderItems: orderItems,
		Status:     "Unpaid",
	}
}

func CartToOrderItemDBModelRequest(r CartItemToOrderItemRequest) OrderItemDBModel {
	return OrderItemDBModel{
		ProductID: r.ProductID,
		Price:     r.Price,
		Quantity:  r.Quantity,
	}
}

type OrderItemResponse struct {
	ID        int
	ProductID int
	Price     float64
	Quantity  int8
}

type OrderResponse struct {
	ID         int
	UserID     int
	OrderItems []OrderItemResponse
	Total      float64
}

func OrderItemResponseFromDBModel(o OrderItemDBModel) OrderItemResponse {
	return OrderItemResponse{
		ID:        o.ID,
		ProductID: o.ProductID,
		Price:     o.Price,
		Quantity:  o.Quantity,
	}
}

func OrderResponseFromDBModel(o OrderDBModel) OrderResponse {
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 20)
	return string(bytes), err
}
