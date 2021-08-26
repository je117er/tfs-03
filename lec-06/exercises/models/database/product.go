package database

import (
	"errors"
	"gorm.io/gorm"
)

type ProductDBModel struct {
	ID          int    `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Price       float64
	Quantity    int8
	Status      string
	CartItems   []CartItemDBModel  `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
	OrderItems  []OrderItemDBModel `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}

func (ProductDBModel) TableName() string {
	return "products"
}

type ProductStorage struct {
	db *gorm.DB
}

func (s ProductStorage) All() ([]ProductDBModel, error) {
	var products []ProductDBModel

	result := s.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (s ProductStorage) ByID(id int) (ProductDBModel, error) {
	var product ProductDBModel

	result := s.db.First(&product, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ProductDBModel{}, ErrProductNotFound
		}
		return ProductDBModel{}, result.Error
	}
	return product, nil
}

func (s ProductStorage) Add(product ProductDBModel) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Omit("CartItems", "OrderItems").Create(&product)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}

func (s ProductStorage) Update(product ProductDBModel) error {
	result := s.db.Save(product)
	return result.Error
}

func (s ProductStorage) Delete(id int) error {
	result := s.db.Delete(&ProductDBModel{}, id)
	return result.Error
}
