package main

import (
	"fmt"
)

// reader := bufio.NewReader(os.Stdin)
// myname, _ := reader.ReadString('\n')
// fmt.Printf("name %s", myname)
//const input = "Now is the winter of our discontent,Made glorious summer by this sun of York."
// reader := bufio.NewReadWriter(strings.NewReader(input))
// reader.Peek('\n')
// fmt.Print("reader", reader)

func main() {
	var n int
	fmt.Scan(&n)
	arr := set(n)
	fmt.Println(minMoves(arr, len(arr)))
}
func minMoves(arr []int, n int) int {
	expectedItem := n

	for i := n - 1; i >= 0; i-- {

		if arr[i] == expectedItem {
			expectedItem--
		}

	}

	return expectedItem

}
func set(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	return a
}
