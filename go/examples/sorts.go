package main

import (
	"fmt"
	"sort"
)

func main() {
	var a []int = []int{9, 4, 6, 2, 6, 0, 5}
	sort.Ints(a)
	fmt.Println(a)

	var b []string = []string{"d", "v", "g", "e"}
	sort.Strings(b)
	fmt.Println(b)

	s := sort.IntsAreSorted(a)
	fmt.Println("sorted", s)
}
