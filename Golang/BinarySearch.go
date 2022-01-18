package main

import (
	"fmt"
)

// Binary Search -> return index of target number otherwise -1
// iterative binary search in go

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	const target = 2

	// result := BinarySearch(array, target)
	result := BinarySearchRec(array, 0, len(array)-1, target)
	fmt.Println(result)
}

func BinarySearch(array []int, target int) int {
	left, right := 0, len(array)-1

	for left <= right {
		mid := (left + right) / 2
		value := array[mid]

		if value == target {
			return mid
		} else if value > target {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	return -1
}

// Recursive binary search in go
func BinarySearchRec(array []int, left, right, target int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	value := array[mid]

	if value == target {
		return mid
	} else if value > target {
		return BinarySearchRec(array, left, mid-1, target)
	} else {
		return BinarySearchRec(array, mid+1, right, target)
	}
}
