package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	Price       float64
	Quantity    uint8
	CartItems   []CartItem `gorm:"constraint:OnDelete:CASCADE;"`
	OrderItems  []OrderItem
}
