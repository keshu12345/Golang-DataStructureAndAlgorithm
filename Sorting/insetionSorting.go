// TC O(n^2)
//SC O(n)

package main

import (
	"fmt"
)

func insertionSort(arr []int, n int) {
	result := make([]int, n)

	result[0] = arr[0]
	for i := 0; i < n; i++ {
		j := i - 1
		for ; j >= 0; j++ {
			if result[j] > arr[i] {
				result[j+1] = result[j]
			} else {
				break
			}
		}
		result[j+1] = arr[i]
	}
}
func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	insertionSort(arr, len(arr))
}
