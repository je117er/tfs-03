// Main program takes in multiple files as input
// and returns the number of primes and their list
// within the files if possible
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isPrime(n int64) bool {
	if n == 2 {
		return true
	} else if n <= 1 || n%2 == 0 {
		return false
	}
	var i int64 = 3
	for ; i < int64(math.Sqrt(float64(n)))+1; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countAndPrintPrimes(f *os.File) {

	input := bufio.NewScanner(f)
	count := 0
	var listOfPrimes []int64

	for input.Scan() {
		parts := strings.Split(input.Text(), " ")
		for i := range parts {
			if el, err := strconv.ParseInt(parts[i], 10, 64); err != nil {
				continue
			} else if isPrime(el) {
				count++
				listOfPrimes = append(listOfPrimes, el)
			}
		}
	}

	fmt.Printf("Number of primes is %v\n", count)
	if count != 0 {
		fmt.Println("List of number of primes:\n")
		fmt.Println(listOfPrimes)
	}
}

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("No file was found.")
		return
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "finder: %v\n", err)
			}
			fmt.Printf("Results for file %s:\n", arg)
			fmt.Println("")
			countAndPrintPrimes(f)
		}
	}
}
