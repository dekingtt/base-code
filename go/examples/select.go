package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func(c chan string) {
		time.Sleep(time.Second)
		c <- "one"
	}(c1)

	go func(c chan string) {
		time.Sleep(time.Second * 2)
		c <- "two"
	}(c2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
