package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"./arithmetic"
)

// checks for possible exceptions during runtime
func check(err error) {
	if err != nil {
		fmt.Println("An exception has occurred. Please restart the program.")
		os.Exit(1)
	}
}

// conv converts a string to a float64 number
func conv(arr []string, index int) float64 {
	res, err := strconv.ParseFloat(arr[index], 64)
	check(err)
	return res
}

// driven program
func main() {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter method: ")
	reader.Scan()
	method := strings.TrimSpace(reader.Text())

	// regex to removes all spaces surrounding the line
	//re := regexp.MustCompile(`\S+`)

	fmt.Print("Enter x and y on the same line: ")
	var x, y float64
	_, err := fmt.Scan(&x, &y)
	check(err)

	/*
		reader.Scan()
		// reads all valid strings into a slice
		parts := re.FindAllString(reader.Text(), -1)

		// converts strings to float64
		x := conv(parts, 0)
		y := conv(parts, 1)
	*/

	fmt.Println("Result: ")
	// method execution starts there
	switch strings.TrimSpace(method) {
	case "add":
		fmt.Println(arithmetic.Add(x, y))
	case "subtract":
		fmt.Println(arithmetic.Subtract(x, y))
	case "mult":
		fmt.Println(arithmetic.Mult(x, y))
	case "div":
		fmt.Println(arithmetic.Div(x, y))
	default:
		fmt.Println("Method not supported.")
	}
}
