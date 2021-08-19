package model

type Cart struct {
	UserID    uint64 `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	CartItems []CartItem
	Status    string
	Total     float64 `gorm:"total_amount"`
}
