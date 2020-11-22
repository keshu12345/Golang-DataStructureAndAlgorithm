// Time Compexity O(2^n)
//Space Compexity O(lenth of the array)==O(k)

package main

import (
	"fmt"
)

func returnSubsequence(arr []int, ci int, reArr []int) {
	if ci >= len(arr) {
		fmt.Println(reArr)
		return

	}

	returnSubsequence(arr, ci+1, reArr)
	reArr = append(reArr, arr[ci])
	returnSubsequence(arr, ci+1, reArr)
}

func main() {
	slice := []int{1, 2, 3}
	reSlice := []int{}

	returnSubsequence(slice, 0, reSlice)
}
