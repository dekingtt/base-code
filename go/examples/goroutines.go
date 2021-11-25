package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println("from", from)
	}
}

func main() {
	f("direct")

	go f("goroutin")

	go func(msg string) {
		fmt.Println(msg)
	}("hello world!")

	time.Sleep(time.Second)
	fmt.Println("done")
}
