package main

import (
	"fmt"
)

func reverseBit(n int) []int {
	power := uint32(31)
	fmt.Println(power)
	slice := []int{}
	for n > 0 {

		slice = append(slice, n%2)
		n = n / 2
	}
	return slice
}

func main() {

	fmt.Println(reverseBit(6))
}
