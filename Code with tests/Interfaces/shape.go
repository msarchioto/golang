package main

import (
	"fmt"
	"math"
)

// Shape is implemented by anything that can tell us its Area
type Shape interface {
	Area() float64
}

// Rectangle implementation of Shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Area function for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Circle implementation of Shape
type Circle struct {
	Radius float64
}

// Area returns the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle implementation of Shape
type Triangle struct {
	Base   float64
	Height float64
}

// Area function for Triangle
func (c Triangle) Area() float64 {
	return 0.5 * (c.Base * c.Height)
}

func main() {
	fmt.Println("Test for interfaces and structs")
}
