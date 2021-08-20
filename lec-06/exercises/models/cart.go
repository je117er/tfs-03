package models

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model
	UserID    uint
	CartItems []CartItem `gorm:"constraint:OnDelete:CASCADE;"`
	Status    string
	Total     float64 `gorm:"total_amount"`
}
