// Time Complexity O(log n)
// Binary search applicable only in sorted array

package main

import (
	"fmt"
)

func binarySearchIterative(arr []int, key int) int {
	n := len(arr)
	l := 0
	r := n - 1

	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == key {
			return mid
		} else if arr[mid] > key {

			r = mid - 1

		} else {
			l = mid + 1

		}

	}
	return -1
}
func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(binarySearchIterative(arr, 7))

}
