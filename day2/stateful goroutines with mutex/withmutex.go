package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var readOps uint64
	var writeOps uint64

	var mutex sync.Mutex
	state := make(map[int]int)

	for r := 0; r < 100; r++ {
		go func() {
			for {
				key := rand.Intn(5)
				mutex.Lock()
				_ = state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)

	fmt.Println("readOps:", readOpsFinal)
	fmt.Println("writeOps:", writeOpsFinal)
}
