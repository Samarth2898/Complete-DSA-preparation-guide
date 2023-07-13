package main

import "fmt"

func swap(arr *[]int, i, j int) {
	temp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = temp
}

// In place sorting function using two pointers.
// Time complexity - O(n)
// Space complexity - O(1)
func sortArr(arr *[]int) {
	s := 0
	e := 0
	n := len(*arr) - 1
	for e <= n {
		if (*arr)[e] < 0 {
			swap(arr, s, e)
			e += 1
			s += 1
		} else {
			e += 1
		}
	}

}

func main() {
	arr := []int{-1, -3, 2, -7, 4, 5, -5, -9, 10}
	sortArr(&arr)
	fmt.Println("sorted array: ", arr)
}
