// Main program takes in a file as input
// and check if it contains values from
// a given list
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findOccurences(f *os.File, values []string) {

	input := bufio.NewScanner(f)
	count := make(map[string]int)
	for input.Scan() {
		text := input.Text()
		for i := range values {
			if strings.Contains(text, values[i]) {
				count[text]++
			}
		}
	}
	fileOut, err := os.Create("/tmp/output")
	check(err)
	defer fileOut.Close()

	_, e := fileOut.WriteString("Strings found in file and their occurences: \n\n")
	check(e)
	for k, v := range count {
		s := fmt.Sprintf("String: %v.\nNumber of occurences: %v\n\n", k, v)
		_, err := fileOut.WriteString(s)
		check(err)
	}

}

func main() {
	n := len(os.Args)
	values := os.Args[1:n]
	file := os.Args[n-1]
	f, err := os.Open(file)
	defer f.Close()
	check(err)
	findOccurences(f, values)
}
