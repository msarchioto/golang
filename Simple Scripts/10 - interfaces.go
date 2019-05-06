// interfaces implement polymorphism
package main

import "fmt"
import "math"

func main() {

	rect1 := Rectangle{20, 50}
	circ := Circle{4}
	fmt.Println(getArea(rect1))
	fmt.Println(getArea(circ))

}

// struct def
type Shape interface {
	area() float64
}
type Rectangle struct {
	height float64
	width  float64
}
type Circle struct {
	radius float64
}

// method area. func (RECEIVER)
func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func getArea(shape Shape) float64 {
	return shape.area()
}
