package main

import "fmt"

func val() (int, int) {
	return 8, 9
}

func main() {
	a, b := val()
	fmt.Println(a, b)

	_, c := val()
	fmt.Println(c)
}
