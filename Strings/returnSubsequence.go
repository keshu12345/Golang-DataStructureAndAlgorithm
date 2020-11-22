// step 1 break the string  xyz to x and yz
// step 2 call recursive call to "yz"{"","z","y","yz"}
// step 3 append  x  in step 2  result  become [ "", x , y , xy , z , xz , yz , xyz ]

package main

import (
	"fmt"
)

func returnSubsequence(input string) []string {
	if len(input) == 0 {
		var base []string
		base = append(base, "")
		return base
	}

	firstCh := input[0:1]

	restOftheString := input[1:]
	restSlice := []string{}
	restSlice = returnSubsequence(restOftheString)

	resultSlice := []string{}

	for _, val := range restSlice {
		resultSlice = append(resultSlice, val)
		resultSlice = append(resultSlice, firstCh+val)

	}
	return resultSlice

}

func main() {

	slice := returnSubsequence("xyz")

	fmt.Println(slice)
	fmt.Println(len(slice))
}
