package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

func (t triangle) getArea() float64 {
	return (0.5 * t.base * t.height)
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return (s.sideLength * s.sideLength)
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println("Total area for shape is:", s.getArea())

}

func main() {
	fmt.Println("Starting")
	t1 := triangle{base: 10, height: 10}
	s1 := square{sideLength: 10}
	fmt.Print("Triangle: ")
	printArea(t1)
	fmt.Print("Square: ")
	printArea(s1)

}
