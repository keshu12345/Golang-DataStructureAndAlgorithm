// Linealy Search the each element it may may sorted and unsorted array or linkedlist
// Time complexity O(n)

package main

import (
	"fmt"
)

func LinearSearch(arr []int, target int) int {

	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}
func main() {
	arr := []int{2, 7, 1, 5, 9, 0, 15}

	fmt.Println(LinearSearch(arr, 15))
}
