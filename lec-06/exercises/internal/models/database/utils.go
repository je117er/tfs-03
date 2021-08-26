package database

import (
	"errors"
)

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrProductNotFound  = errors.New("product not found")
	ErrCartNotFound     = errors.New("cart not found")
	ErrCartItemNotFound = errors.New("cart item not found")
	ErrOrderNotFound    = errors.New("order not found")
)
