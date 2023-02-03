package main

import "fmt"

type GeometricShape interface {
	getPerimeter() float64
}

type Rectangle struct {
	base   float64
	height float64
}
type Square struct {
	side float64
}

func (s Square) getPerimeter() float64 {
	return s.side * s.side
}

func (r Rectangle) getPerimeter() float64 {
	return (r.base * r.height) / 2
}

func getPerimeter(gs1 GeometricShape) float64 {
	return gs1.getPerimeter()
}

func main() {

	r1 := Rectangle{base: 2, height: 1}

	s := Square{2}

	fmt.Printf(" Perimeter of rectangle  : %.0f\n", getPerimeter(r1))
	fmt.Printf(" Perimeter of square : %.1f\n", getPerimeter(s))

}
