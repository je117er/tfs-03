package database

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
