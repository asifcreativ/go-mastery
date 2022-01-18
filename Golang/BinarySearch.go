package main

import (
	"fmt"
)

// Binary Search -> return index of target number otherwise -1
// iterative binary search in go

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	const target = 6

	result := BinarySearch(array, target)
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
