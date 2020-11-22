//Step 1: try to print all bits
// step 2: find the first bit from left
// step 3: make the mask and put that bit in to the result bits int0 right most place

package main

import "fmt"

func reverseBit(n int) {

	flage := false
	var rev uint32
	fmt.Println(rev)
	var j uint32

	for i := 31; i >= 0; i-- {
		mask := (1 << i)
		if flage == true {
			if (mask & n) != 0 {
				fmt.Print(1)
				var maskr uint32
				maskr = 1 << j
				rev = rev | maskr
			} else {
				fmt.Print(0)
			}
			j++
		} else {
			if (mask & n) != 0 {
				flage = true
				fmt.Print(1)
				var maskrev uint32
				maskrev = 1 << j
				rev = rev | maskrev
				j++

			} else {

			}
		}
	}
	fmt.Println()
	fmt.Println("rev", rev)
}

func main() {

	reverseBit(16)
}
