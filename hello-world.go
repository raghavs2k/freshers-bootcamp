package main

import (
	"fmt"
	"math"
	"math/rand"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
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
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}

	fmt.Println("sum2:", sum2)
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

}
