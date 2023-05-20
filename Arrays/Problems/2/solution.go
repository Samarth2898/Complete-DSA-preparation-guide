package main

import "fmt"

type minMaxStruct struct {
	min int
	max int
}

// iterative approach
// 2(n-2)+1 number of comparisons.
// time complexity: O(n)
// space complexity: O(1)
func findMinMax(arr []int) minMaxStruct {
	var minMax minMaxStruct
	n := len(arr)
	if n == 1 {
		minMax.min = arr[0]
		minMax.max = arr[0]
		return minMax
	}

	if arr[0] > arr[1] {
		minMax.max = arr[0]
		minMax.min = arr[1]
	} else {
		minMax.max = arr[1]
		minMax.min = arr[0]
	}

	for i := 2; i < n; i++ {
		if arr[i] > minMax.max {
			minMax.max = arr[i]
		} else if arr[i] < minMax.min {
			minMax.min = arr[i]
		}
	}

	return minMax
}

func main() {
	arr := []int{2, 1, 5, 4, 10, 9, 6}
	minMax := findMinMax(arr)
	fmt.Println("minimum element: ", minMax.min)
	fmt.Println("maximum element: ", minMax.max)
}
