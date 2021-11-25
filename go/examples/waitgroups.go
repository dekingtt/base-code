package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int) {
	fmt.Println("Worker", i, "starting")
	time.Sleep(time.Second)
	fmt.Println("Worker", i, "done")
}

func main() {
	var wg sync.WaitGroup

	const wcount = 5
	for i := 1; i <= wcount; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			worker(id)
		}(i)
	}

	wg.Wait()
}
