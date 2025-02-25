package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("My favorite number is", rand.Intn(10))
	var x int = 10
	y := 1000
	var a = true
	fmt.Println(x, y, a)
	fmt.Println("1+1 =", 1+1)
	fmt.Println("go" + "lang")
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
