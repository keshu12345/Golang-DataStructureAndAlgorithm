// Time complexity  BestCase O(n) and worstCase O(n^2)
// Space Complexity O(1)
package main

import (
	"fmt"
)

func bubbleSort(arr []int, n int) {

	for i := 0; i < n; i++ {
		flage := 0
		for j := n - 1; j >= i+1; j-- {
			if arr[j] < arr[j-1] {
				fmt.Println(i)
				arr[j], arr[j-1] = arr[j-1], arr[j]
				flage = 1
			}
		}
		if flage == 0 {
			break
		}
	}
	fmt.Println(arr)
}
func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	bubbleSort(arr, len(arr))
}
