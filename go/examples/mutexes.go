package main

import (
	"fmt"
	"sync"
)

type container struct {
	mu      sync.Mutex
	counter map[string]int
}

func (c *container) inc(s string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter[s]++
}

func main() {
	var wg sync.WaitGroup
	c := container{
		counter: map[string]int{"a": 0, "b": 0},
	}
	doIncrement := func(k string, v int32) {
		for i := int32(0); i < v; i++ {
			c.inc(k)
		}
		wg.Done()
	}
	wg.Add(3)
	doIncrement("a", 1000)
	doIncrement("a", 1000)
	doIncrement("b", 1000)

	fmt.Println(c.counter)
}
