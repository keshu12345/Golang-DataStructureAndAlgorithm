// Time Complexity O(lon n)
//Space Complexity O(1)

package main

import (
	"fmt"
)

func recursiveBinarySearch(arr []int, l, r, target int) int {

	for l <= r {
		mid := int((l + r) / 2)

		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			return recursiveBinarySearch(arr, l, mid-1, target)
		} else {
			return recursiveBinarySearch(arr, mid+1, r, target)
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 5, 6, 7, 8, 10}
	n := len(arr)
	l := 0
	r := n - 1
	fmt.Println(recursiveBinarySearch(arr, l, r, 10))
}
