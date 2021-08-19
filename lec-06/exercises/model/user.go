package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	//ID uint64
	Username string `gorm:"unique_index;not null"`
	Email    string `gorm:"unique_index;not null"`
	Password string `gorm:"not null"`
	Names    string
	Carts    []Cart
}
