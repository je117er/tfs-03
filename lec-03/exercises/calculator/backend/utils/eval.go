package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func precedence(op rune) int {
	switch op {
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}

func evaluate(x, y string, op rune) string {

	a, _ := strconv.ParseFloat(x, 64)
	b, _ := strconv.ParseFloat(y, 64)
	res := 0.0

	switch op {
	case '+':
		res = a + b
	case '-':
		res = b - a
	case '*':
		res = a * b
	case '/':
		res = b / a
	}
	return abs(res, 1e-5)
}

func abs(a float64, eps float64) string {
	round := math.Round(a)
	if math.Abs(a - round) < eps {
		return fmt.Sprintf("%v", round)
	}
	return fmt.Sprintf("%v", a)
}

func Eval(expr string) (string, error) {

	fmt.Println(expr)

	literals := new(Stack)
	ops := new(Stack)

	var literal strings.Builder

	for i, c := range expr {
		if  unicode.IsDigit(c) || c == '.' {
				literal.WriteRune(c)
				if i == len(expr) - 1 {
					literals.Push(literal.String())
				} else {
					continue
				}
		} else if precedence(c) > 0 {
			literals.Push(literal.String())
			literal.Reset()

			for !ops.IsEmpty() && precedence(c) <= precedence(ops.TopValue().(int32)) {
				value := evaluate(literals.Pop().(string), literals.Pop().(string), ops.Pop().(int32))
				literals.Push(value)
			}
			ops.Push(c)

		}
	}

	for !ops.IsEmpty() {
		value := evaluate(literals.Pop().(string), literals.Pop().(string), ops.Pop().(int32))
		literals.Push(value)
	}
	return literals.TopValue().(string), nil
}

/*
func Eval(expr string) (string, error) {

	m := map[string]bool{"+": true, "-": true, "*": true, "/": true}

	literals := new(Stack)
	op := new(Stack)
	for i, char := range expr {
		if i == 0 {
			if precedence(char) > 0 {
				return "", errors.New("invalid expression")
			}
			literals.Push(string(char))
		} else {
			if unicode.IsDigit(char)  ||  char == '.' {
				if literals.IsEmpty() && op.IsEmpty() {
					return string(char), nil
				} else if literals.IsEmpty() && !op.IsEmpty() {
					return "", errors.New("invalid expression")
				} else if !literals.IsEmpty() && op.IsEmpty() {
					top := literals.Pop().(string)
					literals.Push(top + string(char))
				} else if i == len(expr) - 1 {
					val := arithmetic(literals.Pop().(string), literals.Pop().(string), op.Pop().(int32))

				}
				literals.Push(string(char))
			} else if r := fmt.Sprintf("%c", char); m[r] {
				if (op.IsEmpty())
			}
		}

	}
	return "", nil
}

 */



