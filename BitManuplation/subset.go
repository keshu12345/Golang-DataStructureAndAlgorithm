// step-1: take the loop for 1 leftshift length of the array \
// step-2: take for loop till length of the array
// step-3 check the set bit
// step-4 add into the slice

// Time Complexity O(n*2^n)   space complexity O(2^n)

package main

import (
	"fmt"
)

func subset(slice []int) [][]int {
	sliceRes := [][]int{}

	pow := 1 << len(slice)

	for i := 0; i < pow; i++ {
		tepmSlice := []int{}
		for j := 0; j < len(slice); j++ {
			if (i & (1 << j)) != 0 {
				tepmSlice = append(tepmSlice, slice[j])
			}
		}

		sliceRes = append(sliceRes, tepmSlice)

	}
	return sliceRes
}

func main() {
	slice := []int{1, 2, 3}

	fmt.Println(subset(slice))
}
