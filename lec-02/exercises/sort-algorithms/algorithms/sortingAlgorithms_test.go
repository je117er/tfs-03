package algorithms

import (
	"testing"
)

var x = []int{-4, -5, 20, 0, -80, 7, 50, 1, 99, -2984}
var sortedX = []int{-2984, -80, -5, -4, 0, 1, 7, 20, 50, 99}
var lenX = len(x)
var y = []int{0}
var z = []int{}

func TestBubbleSortNormal(t *testing.T) {
	testArr := make([]int, len(x))
	copy(testArr, x)
	BubbleSort(testArr)
	if !equal(testArr, sortedX) {
		t.Errorf(`TestBubbleSortNormal failed: %v`, testArr)
	}
}

func TestQuickSortNormal(t *testing.T) {
	testArr := make([]int, len(x))
	copy(testArr, x)
	QuickSort(testArr, 0, lenX-1)
	if !equal(testArr, sortedX) {
		t.Errorf(`TestQuickSortNormal failed: %v`, testArr)
	}
}

func TestMergeSortNormal(t *testing.T) {
	testArr := make([]int, len(x))
	copy(testArr, x)
	MergeSort(testArr, 0, lenX-1)
	if !equal(testArr, sortedX) {
		t.Errorf(`TestMergeSortNormal failed: %v`, testArr)
	}
}

func TestBubbleSortSingle(t *testing.T) {
	testArr := make([]int, len(y))
	copy(testArr, y)
	BubbleSort(testArr)
	if !equal(testArr, y) {
		t.Errorf(`TestBubbleSortSingle failed: %v`, testArr)
	}
}

func TestQuickSortSingle(t *testing.T) {
	testArr := make([]int, len(y))
	copy(testArr, y)
	QuickSort(testArr, 0, len(y)-1)
	if !equal(testArr, y) {
		t.Errorf(`TestQuickSortNormal failed: %v`, testArr)
	}
}

func TestMergeSortSingle(t *testing.T) {
	testArr := make([]int, len(y))
	copy(testArr, y)
	MergeSort(testArr, 0, len(y)-1)
	if !equal(testArr, y) {
		t.Errorf(`TestMergeSortNormal failed: %v`, testArr)
	}
}

func TestBubbleSortEmpty(t *testing.T) {
	testArr := make([]int, len(z))
	copy(testArr, z)
	BubbleSort(testArr)
	if !equal(testArr, z) {
		t.Errorf(`TestBubbleSortSingle failed: %v`, testArr)
	}
}

func TestQuickSortEmpty(t *testing.T) {
	testArr := make([]int, len(z))
	copy(testArr, z)
	QuickSort(testArr, 0, len(z)-1)
	if !equal(testArr, z) {
		t.Errorf(`TestQuickSortNormal failed: %v`, testArr)
	}
}

func TestMergeSortEmpty(t *testing.T) {
	testArr := make([]int, len(z))
	copy(testArr, z)
	MergeSort(testArr, 0, len(z)-1)
	if !equal(testArr, z) {
		t.Errorf(`TestMergeSortNormal failed: %v`, testArr)
	}
}
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
