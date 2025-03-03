package main

import (
	"fmt"
	"sync"
)

func frequencyWorker(receiver <-chan string, result chan<- map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	for word := range receiver {
		result <- frequency(word)
	}
}

func frequency(word string) map[string]int {
	freq := make(map[string]int)
	for _, char := range word {
		freq[string(char)]++
	}
	return freq
}

func main() {
	arr := []string{"quick", "brown", "fox", "lazy", "dog"}

	receiver := make(chan string, len(arr))
	result := make(chan map[string]int, len(arr))

	var wg sync.WaitGroup

	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go frequencyWorker(receiver, result, &wg)
	}

	for _, word := range arr {
		receiver <- word
	}
	close(receiver)

	wg.Wait()
	close(result)

	finalFreq := make(map[string]int)
	for res := range result {
		for char, count := range res {
			finalFreq[char] += count
		}
	}

	fmt.Println(finalFreq)
}
