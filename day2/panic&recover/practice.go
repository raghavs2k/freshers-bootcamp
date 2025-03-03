package main

import (
	"fmt"
	"math"
)

func SquareRoot(num int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error recovered:", r)
		}
		if num < 0 {
			panic("Number cannot be less than 0")
		}
		ans := math.Sqrt(float64(num))
		fmt.Printf("Square root of %d: %f\n", num, ans)

	}()
}
func main() {

	SquareRoot(36)
	SquareRoot(-36)

}
