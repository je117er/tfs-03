package arithmetic

import (
	"testing"
)

const (
	x = .4234
	y = 983241235.0
	z = 0.0
)

func TestAdd(t *testing.T) {
	if Add(x, y) != 983241235.4234 {
		t.Errorf(`Add(%v + %v) = false`, x, y)
	}
}

func TestSubtract(t *testing.T) {
	if Subtract(x, y) != -983241234.5766 {
		t.Errorf(`Subtract(%v - %v) = false`, x, y)
	}
}

func TestMult(t *testing.T) {
	if Mult(x, y) != 416304338.899 {
		t.Errorf(`Mult(%v * %v) = false`, x, y)
	}
}

func TestDiv(t *testing.T) {
	val, _ := Div(x, y)
	if val != 4.3061660244548225e-10 {
		t.Errorf(`Div(%v / %v) = false`, x, y)
	}
}

func TestDivByZero(t *testing.T) {
	_, err := Div(x, z)
	if err == nil {
		t.Errorf(`Div(%v / %v) = false. Division by zero expected.`, x, z)
	}
}
