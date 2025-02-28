package main

import (
	"fmt"
	"time"
)

func q(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println(name, ":", i)
	}
}
func main() {

	q("hey")
	go q("hey2")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

	messages := make(chan string)

	go func() { messages <- "ping" }()

	e := <-messages
	fmt.Println(e)

	messages2 := make(chan string, 2)
	messages2 <- "buffered"
	messages2 <- "channel"

	fmt.Println(<-messages2)
	fmt.Println(<-messages2)

}
