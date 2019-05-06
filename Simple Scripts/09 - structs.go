package main

import "fmt"

func main() {

	rect1 := Rectangle{leftX: 0, topY: 50, height: 10, width: 10}
	// this also works if you know the data order
	// rect1 := Rectangle{0, 50, 10, 10}

	fmt.Println("Rectangle is", rect1.width, "wide")
	fmt.Println("Rectangle area =", rect1.area())

}

// struct def
type Rectangle struct {
	leftX  float64
	topY   float64
	height float64
	width  float64
}

// method area. func (RECEIVER)
func (rect *Rectangle) area() float64 {
	return rect.width * rect.height
}
