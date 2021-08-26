package database

import (
	"errors"
	"gorm.io/gorm"
)

type CartDBModel struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	CartItems []CartItemDBModel `gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE"`
	Status    string
	Total     float64 `gorm:"total_amount"`
}

func (CartDBModel) TableName() string {
	return "carts"
}

type CartStorage struct {
	db *gorm.DB
}

func (s CartStorage) ByID(id int) (CartDBModel, error) {
	var cart CartDBModel
	result := s.db.Preload("CartItems").First(&cart, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return CartDBModel{}, ErrCartNotFound
		}
		return CartDBModel{}, result.Error
	}
	return cart, nil
}

func (s CartStorage) Add(cart CartDBModel) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Omit("CartItems").Create(&cart)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}

func (s CartStorage) Update(cart CartDBModel) error {
	result := s.db.Save(cart)
	return result.Error
}

func (s CartStorage) Delete(id int) error {
	result := s.db.Delete(&CartDBModel{}, id)
	return result.Error
}
