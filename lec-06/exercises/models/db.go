package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrProductNotFound  = errors.New("product not found")
	ErrCartNotFound     = errors.New("cart not found")
	ErrCartItemNotFound = errors.New("cart item not found")
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

type CartItemDBModel struct {
	ID        int `gorm:"primaryKey"`
	ProductID int
	CartID    int
	Price     float64
	Quantity  int8
}

func (CartItemDBModel) TableName() string {
	return "cart_items"
}

type OrderItemDBModel struct {
	ID        int `gorm:"primaryKey"`
	ProductID int
	OrderID   int
	Price     float64
	Quantity  int8
}

func (OrderItemDBModel) TableName() string {
	return "order_items"
}

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

type CartStorage struct {
	db *gorm.DB
}

func (s CartStorage) ByID(id int) (CartDBModel, error) {
	var cart CartDBModel
	result := s.db.First(&cart, id)
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
