package main

import (
	"fmt"
	"sort"
)

// Sorting the array and returning the element at (K-1)th index.
// Time complexity - O(n log n)
// Space complexity - O(1)

func getKthSmallestElement(arr []int, k int) int {
	sort.Ints(arr)
	return arr[k-1]
}

func main() {
	arr := []int{4, 5, 10, 30, 78, 97, 8}
	k := 3
	kthSmallest := getKthSmallestElement(arr, k)
	fmt.Println("Kth smallest element in the given unsorted array: ", kthSmallest)
}
