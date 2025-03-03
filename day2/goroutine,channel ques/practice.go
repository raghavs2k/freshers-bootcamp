package main

import (
	"fmt"
	"sync"
)

type fact struct {
	num    int
	result int
}

func factorialWorker(jobs <-chan int, results chan<- fact, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range jobs {
		results <- fact{num: num, result: factorial(num)}
	}
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	arr := []int{5, 7, 10, 2}
	numWorkers := 3

	jobs := make(chan int, len(arr))
	results := make(chan fact, len(arr))

	var wg sync.WaitGroup

	for _, num := range arr {
		jobs <- num
	}
	close(jobs)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go factorialWorker(jobs, results, &wg)
	}

	wg.Wait()
	close(results)

	for res := range results {
		fmt.Printf("Factorial of %d is %d\n", res.num, res.result)
	}
}
