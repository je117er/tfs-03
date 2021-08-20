package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"unique_index;not null"`
	Email        string `gorm:"not null"`
	PasswordHash string `gorm:"not null"`
	Names        string
	Carts        []Cart  `gorm:"constraint:OnDelete:CASCADE;"`
	Orders       []Order `gorm:"constraint:OnDelete:CASCADE;"`
}
