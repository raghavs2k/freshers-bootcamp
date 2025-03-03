package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func collectRating(ratings chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	rating := rand.Intn(5) + 1
	ratings <- rating
}

func main() {

	numStudents := 200
	ratings := make(chan int, numStudents)

	var wg sync.WaitGroup

	for i := 0; i < numStudents; i++ {
		wg.Add(1)
		go collectRating(ratings, &wg)
	}

	wg.Wait()
	close(ratings)

	total := 0
	count := 0

	for rating := range ratings {
		total += rating
		count++
	}

	averageRating := float64(total) / float64(count)
	fmt.Printf("Average Teacher Rating: %.2f\n", averageRating)
}
