package main

import "fmt"

func main() {
	shouldRecover := false
	dontPanic(shouldRecover)
	fmt.Println("magical")
}
