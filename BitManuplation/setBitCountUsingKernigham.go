// step-1: calculate the most resent right set bit
// step-2:  rsbm is subtract number
// step-3: count++
// while number become zero

// Timecomplexity O(1)  same as space complexity

package main

import (
	"fmt"
)

func bitCount(n int) int {
	counter := 0
	for n != 0 {

		rsbm := n & (-n)
		n = n - rsbm
		counter++
	}
	return counter

}

func main() {

	fmt.Println(bitCount(13))
}
