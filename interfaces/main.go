package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

func main() {
	t := triangle{base: 22, height: 11}
	s := square{side: 12}
	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println("thae area is ", s.getArea())
}

func (t triangle) getArea() float64 {
	a := (0.5 * t.base * t.height)
	return a
}
func (s square) getArea() float64 {
	a := s.side * s.side
	return a
}
