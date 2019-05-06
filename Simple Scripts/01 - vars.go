package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Printf("hello, world\n")

	// var declaration
	var age int = 40
	var favNumber float64 = 1.6180339

	fmt.Println(age, favNumber)

	randNum := 1

	var numOne = 1.000
	var num99 = .9999

	fmt.Println("Variable Types")
	fmt.Println(reflect.TypeOf(randNum))
	fmt.Println(reflect.TypeOf(numOne))
	fmt.Println(reflect.TypeOf(num99))

}
