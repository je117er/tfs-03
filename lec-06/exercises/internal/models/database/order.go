package database

import (
	"errors"
	"gorm.io/gorm"
)

type OrderDBModel struct {
	ID         int `gorm:"primaryKey"`
	UserID     int
	OrderItems []OrderItemDBModel `gorm:"foreignKey:OrderID;constraint:OnDelete:No ACTION"`
	Status     string
	Total      float64
}

func (OrderDBModel) TableName() string {
	return "orders"
}

type OrderStorage struct {
	db *gorm.DB
}

func (s OrderStorage) All() ([]OrderDBModel, error) {
	var orders []OrderDBModel
	result := s.db.Preload("OrderItems").Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (s OrderStorage) ByID(id int) (OrderDBModel, error) {
	var order OrderDBModel
	result := s.db.Preload("OrderItems").First(&order, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return OrderDBModel{}, ErrOrderNotFound
		}
		return OrderDBModel{}, result.Error
	}
	return order, nil
}

func (s OrderStorage) Add(order OrderDBModel) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Omit("OrderItems").Create(order)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}
