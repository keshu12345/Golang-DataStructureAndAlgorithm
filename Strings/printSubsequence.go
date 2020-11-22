// step 1 take input string  "xyz" and output string " "
// step 2  string "xyz" there is two choice either include x or exclude the x into the output string
// step 3 recursive call
// Time Complexity O(2^n)

package main

import (
	"fmt"
)

func printSubsequence(input, output string) {
	if len(input) == 0 {
		fmt.Println(output)
		return
	}

	printSubsequence(input[1:], output)
	printSubsequence(input[1:], output+input[0:1])

}

func main() {
	printSubsequence("xyz", "")
}
