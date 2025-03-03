package main

import (
	"fmt"
	"sync"
)

type container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *container) inc(a string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[a]++
}

func main() {
	c := container{
		counters: map[string]int{"x": 0, "y": 0},
	}
	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()

	}

	wg.Add(3)
	go doIncrement("x", 1000)
	go doIncrement("y", 2000)
	go doIncrement("x", 3000)

	wg.Wait()
	fmt.Println(c.counters)
}
