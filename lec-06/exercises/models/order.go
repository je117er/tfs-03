package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	UserID     uint
	OrderItems []OrderItem `gorm:"constraint:OnDelete:CASCADE;"`
	Status     string
	Total      float64
}
