package main

import (
	"fmt"
)

func selectionSort(arr []int, n int) []int {
	var min int
	index := 0

	for i := 0; i < n; i++ {
		min = arr[i]
		for j := i + 1; j < n; j++ {
			if arr[i] < min {
				min = arr[i]
				index = j
			}
		}
		arr[i], arr[index] = arr[index], arr[i]
	}

}
func main() {
	var n int

	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(selectionSort(arr, len(arr)))
}
