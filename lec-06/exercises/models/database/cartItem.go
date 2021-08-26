package database

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
