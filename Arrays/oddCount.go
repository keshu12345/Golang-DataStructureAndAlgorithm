// find the number occuring odd number of times in an array given that exactly one number odd number of times?
// Naive approach
// time Complexity O(n^2)

package main

import (
	"fmt"
)

func oddOccuranceNaiveApproach(arr []int) int {
	result := 0
	count := 0

	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				count++
			}
		}
		count++
		if (count % 2) != 0 {
			result = arr[i]
			break
		}
		count = 1

	}
	return result
}
func main() {
	arr := []int{1, 1, 4, 3, 4, 3, 4, 5, 5}

	fmt.Println(oddOccuranceNaiveApproach(arr))
}
