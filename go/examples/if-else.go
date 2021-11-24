package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if n := 9; n < 0 {
		fmt.Println("negtive")
	} else if n < 10 {
		fmt.Println("one digit")
	} else {
		fmt.Println("mutiple digit")
	}
}
