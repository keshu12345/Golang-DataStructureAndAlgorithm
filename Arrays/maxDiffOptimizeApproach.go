// Time Complixity O(n)
// Space COmpexity O(size of the pair)

package main

import (
	"fmt"
	"math"
)

func maxDiffOptimal(arr []int) ([2]int, int) {

	pairs := [2]int{}
	target := math.MinInt64
	min_so_far := arr[0]
	max_diff_so_far := arr[1] - arr[0]
	current_diff := arr[1] - arr[0]
	for index := 1; index < len(arr); index++ {
		if arr[index] < min_so_far {
			min_so_far = arr[index]
		} else {
			current_diff = arr[index] - min_so_far
			if current_diff > max_diff_so_far {
				max_diff_so_far = current_diff
				target = max_diff_so_far
				pairs[0] = min_so_far
				pairs[1] = arr[index]

			}
		}
	}
	return pairs, target
}

func main() {

	arr := []int{3, 7, 8, 5, 1, 2, 10, 15, 0, 26, 1, 2, 3, 87}
	fmt.Println(maxDiffOptimal(arr))
}
