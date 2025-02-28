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
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
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

	done := make(chan bool, 1)
	go worker(done)

	<-done

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-c1:
			fmt.Println("Recieved", msg)
		case msg2 := <-c2:
			fmt.Println("Recieved", msg2)
		}
	}

	signal := make(chan string)
	rag := make(chan string)

	select {
	case msg := <-signal:
		fmt.Println("recieved message", msg)
	default:
		fmt.Println("nothing received")
	}

	msg := "hey"
	select {
	case signal <- msg:
		fmt.Println("message sent", signal)
	default:
		fmt.Println("nothing sent")
	}

	select {
	case msg := <-signal:
		fmt.Println("signal recieved", msg)
	case msg2 := <-rag:
		fmt.Println("rag recieved", msg2)
	default:
		fmt.Println("no activity")

	}

	jobs := make(chan int, 5)
	done2 := make(chan bool)

	go func() {
		for {
			j, full := <-jobs
			if full {
				fmt.Println("Receiving:", j)
			} else {
				fmt.Println("All Recieved")
				done2 <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("Sent Job:", i)
	}
	close(jobs)
	fmt.Println("Sent all jobs")

	<-done2

	_, ok := <-jobs
	fmt.Println("receiving more jobs:", ok)

}
