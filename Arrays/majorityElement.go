// Time Complexity O(nlogn)
//Space COMplexity O(1)
// Given an array A of size n, find an element that occurs more than n/2 times?
package main

import (
	"fmt"
)

func majorityElement(arr []int) int {
	n := (len(arr) / 2)

	for i := 0; i < len(arr)/2; i++ {
		if arr[i] == arr[i+n] {
			return arr[i]
		}
	}

	return -1

}
func main() {

	arr := []int{2, 3, 2, 2, 3, 2, 4, 5, 6}

	fmt.Println(majorityElement(arr))

}
