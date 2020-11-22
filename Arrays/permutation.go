package main

import (
	"fmt"
)

var count int

func permuatation(arr []int, index int) {

	if index == len(arr)-1 {
		fmt.Println(arr)

	}

	for i := index; i < len(arr); i++ {

		arr[index], arr[i] = arr[i], arr[index]
		permuatation(arr, index+1)
		arr[index], arr[i] = arr[i], arr[index]

	}

}

func main() {

	arr := []int{1, 2, 3}

	permuatation(arr, 0)
}
