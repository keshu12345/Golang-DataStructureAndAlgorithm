// find maximum difference between two elements in an array such that larger element appears after the smaller number
// TimeComplexity O(n^2)  //Brute Force
// Space Compexity O(k)   // TOTAL SIZE OF allocated to the array

//Naive Approach
package main

import (
	"fmt"
	"math"
)

func maximumDiff(arr []int) ([2]int, int) {

	nums := [2]int{}
	maximumDiff := math.MinInt64
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				current_diff := arr[j] - arr[i]
				if current_diff > maximumDiff {
					maximumDiff = current_diff
					nums[0] = arr[i]
					nums[1] = arr[j]
				}
			}
		}
	}

	return nums, maximumDiff

}

func main() {
	arr := []int{-2, 2, 5, 7, 8, 10, 1, 15, 20, 0, 20, 45}

	fmt.Println(maximumDiff(arr))

}
