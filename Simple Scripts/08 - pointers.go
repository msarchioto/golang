package main

import "fmt"

func main() {

	x := 0
	changeXVal(x)
	fmt.Println("x = ", x)

	changeXValNow(&x)
	fmt.Println("x = ", x)

	fmt.Println("Memory address for x =", &x)

	// define pointer
	yPtr := new(int)
	fmt.Println("Memory address for yPtr =", yPtr)
	// change pointer value
	changeYValNow(yPtr)
	fmt.Println("Memory address for yPtr =", *yPtr)

}

func changeYValNow(yPtr *int) {
	*yPtr = 200
}

func changeXValNow(x *int) {
	*x = 2
}

func changeXVal(x int) {
	x = 2
}
