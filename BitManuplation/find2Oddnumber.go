// find the two number that is odd occurance time's
// step-1 take xor of all number
// step-2  find the most set right bit of the given number
// step-3 seprate the those number on the basis of msb

// Time Complexity O(n)
// Space compexity O(1)

// if try to use Hashmap than timecomlexity O(n)  and spcace compexity O(n)  when bit masking most efficient approach to solve this problem

package main

import (
	"fmt"
)

func printTwoNumberOddtimesOccur(slice []int) []int {

	sliceRe := []int{}

	xor := 0
	for _, val := range slice {
		xor = xor ^ val
	}
	msb := xor & (-xor)

	num1 := 0
	num2 := 0
	for _, val := range slice {
		if (val & msb) == 0 {
			num1 = num1 ^ val
		} else {
			num2 = num2 ^ val
		}
	}
	sliceRe = append(sliceRe, num1, num2)

	return sliceRe

}

func main() {

	slice := []int{1, 2, 4, 5, 1, 5, 2, 9, 4, 9, 7, 5}

	fmt.Println(printTwoNumberOddtimesOccur(slice))
}
