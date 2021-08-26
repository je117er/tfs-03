package database

import (
	"gorm.io/gorm"
	"time"
)

type PaymentDBModel struct {
	ID           int `gorm:"primaryKey"`
	OrderID      int
	Amount       float64
	ReceiptEmail string
	PaymentDate  time.Time
}

func (PaymentDBModel) TableName() string {
	return "payments"
}

type PaymentStorage struct {
	db *gorm.DB
}

func (s PaymentStorage) Add(p PaymentDBModel) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&p)
		if result.Error != nil {
			return result.Error
		}
		err := tx.Model(&OrderDBModel{}).Where("ID = ?", p.OrderID).Update("status", "paid").Error
		if err != nil {
			return err
		}
		return nil
	})
}
