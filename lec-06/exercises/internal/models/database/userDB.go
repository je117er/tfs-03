package database

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type UserDBModel struct {
	ID           int    `gorm:"primaryKey"`
	Username     string `gorm:"not null"`
	Email        string `gorm:"not null"`
	PasswordHash string `gorm:"not null"`
	Names        string `gorm:"not null"`
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	Carts        []CartDBModel  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Orders       []OrderDBModel `gorm:"foreignKey:UserID;constraint:OnDelete:NO ACTION"`
}

func (UserDBModel) TableName() string {
	return "users"
}

type UserStorage struct {
	db *gorm.DB
}

func (s UserStorage) ByID(id int) (UserDBModel, error) {
	var user UserDBModel

	result := s.db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return UserDBModel{}, result.Error
		}
		return UserDBModel{}, result.Error
	}

	return user, nil
}
func (s UserStorage) Add(user UserDBModel) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Omit("Carts", "Orders").Create(&user)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}

func (s UserStorage) Update(user UserDBModel) error {
	result := s.db.Save(user)
	return result.Error
}

func (s UserStorage) Delete(id int) error {
	result := s.db.Delete(&UserDBModel{}, id)
	return result.Error
}
