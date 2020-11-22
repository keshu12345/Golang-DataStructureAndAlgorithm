package main

import (
	"fmt"

	"sort"
	"strings"
)

func stringsort(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")

}

func main() {

	s := []string{"tac", "dog", "god", "cat"}

	m := map[string][]int{}

	for key, val := range s {
		str := stringsort(val)
		m[str] = append(m[str], key)

	}
	slice := [][]int{}
	for key, val := range m {
		slice = append(slice, val)
		fmt.Println(key + " ")
		fmt.Println(val)

	}
	fmt.Println(slice)

}
