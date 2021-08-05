package main

import (
	"fmt"

	"./algorithms"
)

func main() {
	arrBubble := []int{2, 4, 3, 0, -1}
	//algorithms.BubbleSort(arr[:])
	left, right := 0, len(arr)-1
	//algorithms.MergeSort(arr, left, right)
	algorithms.QuickSort(arr, left, right)
	fmt.Println(arr)

}
