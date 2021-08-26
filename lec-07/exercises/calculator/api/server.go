package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

type Result struct {
	Result string `json:"result"`
}

func Handler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/api/calc" || req.Method != "GET" {
		log.Println("oh no")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Content-Type", "application/json")

	q := req.URL.Query()
	fmt.Printf("%v\n", q)
	exp := q["exp"][0]
	res, err := Eval(exp)
	if err != nil {
		msg := fmt.Sprintf("Invalid expression: %s\n", exp)
		http.Error(w, msg, http.StatusInternalServerError)
	}
	res1D := &Result{Result: res}
	_ = json.NewEncoder(w).Encode(res1D)
	fmt.Println(res)
	return
}

func Precedence(op rune) int {
	switch op {
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}

func Evaluate(x, y string, op rune) string {

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
	return Abs(res, 1e-5)
}

func Abs(a float64, eps float64) string {
	round := math.Round(a)
	if math.Abs(a-round) < eps {
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
		if unicode.IsDigit(c) || c == '.' {
			literal.WriteRune(c)
			if i == len(expr)-1 {
				literals.Push(literal.String())
			} else {
				continue
			}
		} else if Precedence(c) > 0 {
			literals.Push(literal.String())
			literal.Reset()

			for !ops.IsEmpty() && Precedence(c) <= Precedence(ops.TopValue().(int32)) {
				value := Evaluate(literals.Pop().(string), literals.Pop().(string), ops.Pop().(int32))
				literals.Push(value)
			}
			ops.Push(c)

		} else {
			return "", errors.New("invalid expression")
		}
	}

	for !ops.IsEmpty() {
		value := Evaluate(literals.Pop().(string), literals.Pop().(string), ops.Pop().(int32))
		literals.Push(value)
	}
	return literals.TopValue().(string), nil
}

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *Stack) TopValue() (value interface{}) {
	return s.top.value
}

func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Clear() (value interface{}) {
	if s.size > 0 {
		value, s.top = nil, nil
		s.size = 0
	}
	return nil
}
