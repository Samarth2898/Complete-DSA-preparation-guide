package main

import "fmt"

// iterative approach
// time complexity O(n)
// space complexity O(1)
func reverseArray(arr *[]int) {
	start := 0
	end := len(*arr) - 1
	for start < end {
		temp := (*arr)[start]
		(*arr)[start] = (*arr)[end]
		(*arr)[end] = temp
		start += 1
		end -= 1
	}
}

// recursive approach
// time complexity O(n)
// space complexity O(1)
func reverseArrayRecursive(arr *[]int, start int, end int) {
	if start >= end {
		return
	}
	temp := (*arr)[start]
	(*arr)[start] = (*arr)[end]
	(*arr)[end] = temp

	reverseArrayRecursive(arr, start+1, end-1)
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println("original array: ", arr)
	reverseArray(&arr)
	fmt.Println("reversed array: ", arr)

	arr2 := []int{5, 6, 7, 8}
	fmt.Println("original array approach 2: ", arr2)
	reverseArrayRecursive(&arr2, 0, len(arr2)-1)
	fmt.Println("reversed array approach 2: ", arr2)
}
