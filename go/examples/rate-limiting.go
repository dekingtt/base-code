package main

import (
	"fmt"
	"time"
)

func main() {
	request := make(chan int, 5)
	for i := 0; i < 5; i++ {
		request <- i
	}
	close(request)

	limiter := time.Tick(time.Millisecond * 200)
	for r := range request {
		<-limiter
		fmt.Println("request", r, time.Now())
	}

	burstLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstLimiter <- t
		}

	}()
	burstRequest := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstRequest <- i
	}
	close(burstRequest)
	for r := range burstRequest {
		<-burstLimiter
		fmt.Println("burst request", r, time.Now())
	}
}
