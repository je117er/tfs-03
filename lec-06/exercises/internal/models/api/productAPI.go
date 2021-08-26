package api

import (
	"exercises/internal/models/database"
)

type CreateProductRequest struct {
	Title       string
	Description string
	Price       float64 `validate:"numeric"`
	Quantity    int8    `validate:"numeric"`
}

func ProductDBModelFromCreateRequest(r CreateProductRequest) database.ProductDBModel {
	return database.ProductDBModel{
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

func ProductResponseFromDBModel(p database.ProductDBModel) ProductResponse {
	return ProductResponse{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    p.Quantity,
		Status:      p.Status,
	}
}
