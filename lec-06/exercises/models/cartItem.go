package models

import "github.com/jinzhu/gorm"

type CartItem struct {
	gorm.Model
	ProductID uint
	CartID    uint
	Price     float64
	Quantity  uint8
}
