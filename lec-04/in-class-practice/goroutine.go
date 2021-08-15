package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Starting...")
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 'a'; i <= 'z'; i++ {
			fmt.Printf("%c ", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i <= 26; i++ {
			fmt.Printf("%d ", i)
		}
	}()
	wg.Wait()
	fmt.Println("Ended")
}
