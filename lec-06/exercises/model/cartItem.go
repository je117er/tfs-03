package model

import "github.com/jinzhu/gorm"

type CartItem struct {
	gorm.Model
	ID        uint64
	ProductID uint64
	CartID    uint64
	Price     float64
	Quantity  uint8
}
