package main

import (
	"fmt"
	"math"
)

type square struct {
	length int
}

type rectangle struct {
	length  int
	breadth int
}

type circle struct {
	radius int
}

type i interface {
	area() int
	perimeter() int
}

func (s *square) area() int {
	return s.length * s.length
}

func (r *rectangle) area() int {
	return r.length * r.breadth
}

func (c *circle) area() int {
	return int(math.Pi * float64(c.radius) * float64(c.radius))
}

func (s *square) perimeter() int {
	return 4 * s.length
}

func (r *rectangle) perimeter() int {
	return 2 * (r.length + r.breadth)
}

func (c *circle) perimeter() int {
	return int(2 * math.Pi * float64(c.radius))
}

func print(inter i, ShapeName string) {
	fmt.Printf("%s area : %d\n", ShapeName, inter.area())
	fmt.Printf("%s perimeter: %d\n", ShapeName, inter.perimeter())
}

func main() {
	s := &square{5}
	r := &rectangle{5, 10}
	c := &circle{7}

	print(s, "Square")
	print(r, "Rectangle")
	print(c, "Circle")
}
