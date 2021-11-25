package main

import "fmt"

func sum(nums ...int) int {
	var s int
	for _, v := range nums {
		s += v
	}

	return s
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))

	nums := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	fmt.Println(sum(nums...))
}
