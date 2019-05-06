package main

import (
	"fmt"
)

func main() {

	const pi float64 = 3.1415926

	var (
		varA = 2
		varB = 3
	)

	var myName string = "Mik Sar"
	fmt.Println(len(myName), varA, varB)

	var isOver40 bool = true

	fmt.Printf("%.3f \n", pi)
	// type of var
	fmt.Printf("%T \n", pi)
	fmt.Printf("%t \n", isOver40)

	fmt.Printf("%d \n", 100)
	fmt.Printf("%b \n", 100)
	fmt.Printf("%c \n", 44)
	fmt.Printf("%x \n", 17)
	fmt.Printf("%e \n", pi)
}
