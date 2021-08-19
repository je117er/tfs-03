package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	//ID uint64
	Title       string `gorm:"not null"`
	Description string
	Price       float64
	Quantity    uint8
	CartItems   []CartItem
}
