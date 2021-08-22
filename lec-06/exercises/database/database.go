package database

import (
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
