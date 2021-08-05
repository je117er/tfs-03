// Package arithmetic performs basic calculating functions,
// including addition, subtraction, multiplication and division
package arithmetic

import "errors"

func Add(x, y float64) float64 {
	return x + y
}

func Subtract(x, y float64) float64 {
	return x - y
}

func Mult(x, y float64) float64 {
	return x * y
}

// Returns the quotient of two numbers and an non-nil error if the divisor is 0
func Div(x, y float64) (float64, error) {
	if y == 0.0 {
		return -1, errors.New("Division by 0!")
	}
	return x / y, nil
}
