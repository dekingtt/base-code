package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second)
	<-timer1.C
	fmt.Println("timer 1 fired")
	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()
	timer2.Stop()
	fmt.Println("timer 2 stopped")

	time.Sleep(2 * time.Second)
}
