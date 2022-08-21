package main

import "fmt"

// setting up the Shape interface.
type Shape interface {
	getArea() float64
}

// setting up Triangle struct type.
type Triangle struct {
	Height float64
	Base float64
}

// setting up Square struct type.
type Square struct {
	sideLength float64
}

func main() {
	// adding values to both triangle and square types.
	t := Triangle{4.5, 6.5}
	s := Square{7.33}

	fmt.Println("The area of the triangle is:", printArea(t))
	fmt.Println("The area of the square is:", printArea(s))
}

/*
 * this function takes any struct type that
 * is part of the Shape interface and returns
 * a float64 value resulted from both triangle
 * and square calculations.
 */
func printArea(s Shape) float64 {
	return s.getArea()
}

// calculate the area of the triangle.
func (t Triangle) getArea() float64 {
	return 0.5 * t.Base * t.Height
}

// calculate the area of the square.
func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}