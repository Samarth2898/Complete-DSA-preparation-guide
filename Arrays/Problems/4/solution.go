package main

import "fmt"

func swap(arr *[]int, i, j int) {
	temp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = temp
}

// In place sorting function using three pointers.
// Time complexity - O(n)
// Space complexity - O(1)
func sortArr(arr *[]int) {
	l := 0
	m := 0
	h := len(*arr) - 1
	for m <= h {
		if (*arr)[m] == 0 {
			swap(arr, l, m)
			m += 1
			l += 1
		} else if (*arr)[m] == 1 {
			m += 1
		} else {
			swap(arr, m, h)
			h -= 1
		}
	}
}

func main() {
	arr := []int{1, 0, 2, 1, 1, 2, 0, 0}
	sortArr(&arr)
	fmt.Println("sorted array: ", arr)
}
