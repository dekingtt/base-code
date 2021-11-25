package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "start to work on job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		result <- j * 2
	}
}

func main() {
	const jobnums = 5
	jobs := make(chan int, jobnums)
	result := make(chan int, jobnums)

	for i := 1; i <= jobnums; i++ {
		go worker(i, jobs, result)
	}

	for i := 1; i <= jobnums; i++ {
		jobs <- i
	}
	close(jobs)
	for k := 1; k <= jobnums; k++ {
		<-result
	}

}
