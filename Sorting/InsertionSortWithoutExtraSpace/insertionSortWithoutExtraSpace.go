package main

import (
	"fmt"
)

func InsertionSortMap(arr []int, n int) {

	for i := 0; i < n; i++ {
		temp := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > arr[i] {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = temp
	}
	fmt.Println(arr)
}
func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	InsertionSortMap(arr, len(arr))
}
