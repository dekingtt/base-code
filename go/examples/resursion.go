package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
	var fib func(int) int
	fib = func(i int) int {
		if i < 2 {
			return i
		}
		return fib(i-1) + fib(i-2)
	}

	fmt.Println(fib(7))
}
