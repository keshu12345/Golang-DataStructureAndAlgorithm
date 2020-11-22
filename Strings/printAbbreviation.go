package main

import (
	"fmt"
	"strconv"
)

func printAbbreviation(input, output string, count int) {
	slice := []string{}
	if len(input) == 0 {
		fmt.Println(slice)
	}

	printAbbreviation(input[1:], output, count+1)
	slice = append(slice, output+strconv.Itoa(count)+input[0:1])
	printAbbreviation(input[1:], output+input[0:1], count)

}

func main() {

	printAbbreviation("xyz", "", 0)
}
