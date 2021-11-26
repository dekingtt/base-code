package main

import "fmt"

func mypanic() {
	panic("a problem")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered error", r)
		}
	}()

	mypanic()
	fmt.Println("after panic") // this line will not run because of panic
}
