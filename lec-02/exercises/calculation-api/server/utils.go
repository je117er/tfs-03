package server

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

type operands struct {
	first, second float64
}

type calculation interface {
	add() float64
	subtract() float64

	mult() float64
	div() (float64, error)
}

func (op operands) add() float64 {
	return op.first + op.second
}

func (op operands) subtract() float64 {
	return op.first - op.second
}
func (op operands) mult() float64 {
	return op.first * op.second
}

func (op operands) div() (float64, error) {
	if op.second == 0 {
		return -1, errors.New("Division by 0!")
	}
	return op.first / op.second, nil
}

func getQueryParams(query string) (string, float64, float64, error) {

	methods := map[string]bool{"add": true, "subtract": true, "mult": true, "div": true}

	dict, _ := url.ParseQuery(query)
	if _, ok := methods[dict["method"][0]]; ok {

		first, erf := strconv.ParseFloat(dict["first"][0], 64)
		if erf != nil {
			return "", -1, -1, errors.New("Invalid request")
		}

		second, ers := strconv.ParseFloat(dict["second"][0], 64)
		if ers != nil {
			return "", -1, -1, errors.New("Invalid request")
		}
		return dict["method"][0], first, second, nil
	}
	return "", -1, -1, errors.New("Method not supported")
}

func eval(query string) (string, error) {

	method, first, second, err := getQueryParams(query)
	if err != nil {
		return "", err
	}

	op := operands{first, second}
	var res float64

	switch method {
	case "add":
		res = calculation.add(op)
	case "subtract":
		res = calculation.subtract(op)
	case "mult":
		res = calculation.mult(op)
	default:
		var er error
		res, er = calculation.div(op)
		if er != nil {
			return "", er
		}
	}
	return fmt.Sprintf("%f", res), nil
}
