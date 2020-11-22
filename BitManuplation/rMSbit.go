
// Find the Right most set bit
// step-1 convert numbers into 1's compliment and add 1
//step-2  take and operation to number and 1's compliment+1 you will get result  right most set bit

package main

import (
	"fmt"
)

func rightMostSetBit(n int) int {

	res := n & (-n)
	return res
}

func main() {

	fmt.Println(rightMostSetBit(20))
}
