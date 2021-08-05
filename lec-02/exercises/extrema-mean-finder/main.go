// Main program takes in multiple files as input
// and returns extrema and mean of the numbers
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

func findExtremaMean(f *os.File) {

	input := bufio.NewScanner(f)

	max := math.SmallestNonzeroFloat64
	min := math.MaxFloat64
	var sum float64
	count := 0

	for input.Scan() {
		//	if el, err := strconv.ParseFloat(input.Text(), 64); err != nil {
		parts := strings.Split(input.Text(), " ")
		for i := range parts {
			if el, err := strconv.ParseFloat(parts[i], 64); err != nil {
				continue
			} else {
				count++
				sum += el
				max = math.Max(max, el)
				min = math.Min(min, el)
			}
		}
	}

	// output max
	if max != math.SmallestNonzeroFloat64 {
		fmt.Printf("Max is %v\n", max)
	} else {
		fmt.Println("No max found")
	}

	// output min
	if min != math.MaxFloat64 {
		fmt.Printf("Min is %v\n", min)
	} else {
		fmt.Println("No min found")
	}

	// output mean
	if count == 0 {
		fmt.Println("There's no number in the file.")
	} else {
		fmt.Printf("Mean is %v\n", sum/float64(count))
	}

	fmt.Println("")
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
			findExtremaMean(f)
		}
	}
}
