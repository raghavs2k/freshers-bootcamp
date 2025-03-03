package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops atomic.Uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				ops.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("ops:", ops.Load())

}
