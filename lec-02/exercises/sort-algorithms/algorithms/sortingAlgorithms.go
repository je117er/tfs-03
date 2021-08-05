// Package algorihtms implements three popular sorting algorithms:
// bubble sort, merge sort and quicksort
// Can be only applied to a list of integers
// All algorithms transform an unsorted list of number into
// a non-decreasing sequence of numbers
package algorithms

// Bubble sort - a comparison sort, goes through a list and exchange positions
// of two adjacent elements if they're in the wrong order
// until the whole list is in the right one
// Time complexity is O(n^2) due to two for loops
// Space complexity is O(n) as only the originally array was needed
func BubbleSort(arr []int) {
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// Merge sort is also a comparison sorting algorithm.
// Unlike bubble sort, merge sort uses a divide-and-
// conquerered approach that drastically reduces
// the time complexity compared to bubble sort
//
// A merge sort partitions the initial list into
// single-element sublists and recursively combines
// them into a sorted list
func MergeSort(arr []int, left, right int) {

	if left < right {
		mid := (left + right) / 2
		MergeSort(arr, left, mid)
		MergeSort(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
}

// merge is a helper function that combines
// the sublists from MergeSort
func merge(arr []int, left, mid, right int) {

	//tmp := make([]int, 0)
	tmp := make([]int, len(arr))
	copy(tmp, arr)

	i, j, k := left, mid+1, left

	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			i++
		} else {
			tmp[k] = arr[j]
			j++
		}
		k++
	}
	// Add the remaining non-empty sublist to the merged list
	copy(tmp[k:right+1], arr[j:right+1])
	copy(tmp[k:right+1], arr[i:mid+1])
	copy(arr, tmp)
}

// Quicksort is also a divide and conquer algorithm
// This implementation selects the last element as pivot,
// partitions the list into two left and right subarrays
// consisting of elements less than and greater than
// or equal to the pivot, respectively.

func QuickSort(arr []int, lo, hi int) {
	// checks if lo and hi are natural numbers
	if lo >= 0 && hi >= 0 {
		if lo < hi {
			// finds the index of the pivot
			p := partition(arr, lo, hi)
			// recursively apply quicksort for the left and right side of the array
			QuickSort(arr, lo, p-1)
			QuickSort(arr, p+1, hi)
		}
	}
}

// partition functions returns the index of the pivot
func partition(arr []int, lo, hi int) int {
	// sets the pivot to the last value
	// of the array
	p := arr[hi]
	// sets the pivot index to the start
	// of the array
	i := lo - 1
	// goes thru the array from lo to hi
	for j := lo; j <= hi; j++ {
		// swap the element less than or equal to the pivot
		// with the element at index i and increment i
		if arr[j] <= p {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return i
}
