// TimeComplexity O(n)
//Space Complexity O(1)

package main

import (
	"fmt"
)

func majorityMooreElement(arr []int) int {

	majorityIndex := 0
	count := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[majorityIndex] {
			count++
		} else {
			count--
		}

		if count == 0 {
			majorityIndex = i
			count++
		}
	}
	return arr[majorityIndex]
}
func isMajorityElement(arr []int, size, majorityElement int) bool {
	count := 0
	for i := 0; i < size; i++ {
		if arr[i] == majorityElement {
			count++
		}

	}
	if count > (size / 2) {
		return true
	} else {
		return false
	}

}
func main() {
	arr := []int{2, 1, 1, 1, 3, 4, 5, 6, 1, 1, 1}

	size := len(arr)

	majorityElement := majorityMooreElement(arr)

	fmt.Println(isMajorityElement(arr, size, majorityElement))
}
