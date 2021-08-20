package models

import "github.com/jinzhu/gorm"

type OrderItem struct {
	gorm.Model
	ProductID uint
	OrderID   uint
	Price     float64
	Quantity  uint8
}
