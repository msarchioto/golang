package main

import "fmt"

func main() {

	listNums := []float64{1, 2, 3, 4, 5}

	fmt.Println("Sum :", addAllNumsFromArray(listNums))

	num1, num2 := next2values(5)
	fmt.Println(num1, num2)

	// send undefined number of values to a func
	fmt.Println(subtractThem(1, 2, 3, 4, 5))

	// function inside of a function - "closures"
	num3 := 3
	doubleNum := func() int {
		num3 *= 2
		return num3
	}
	fmt.Println(doubleNum())
	fmt.Println(doubleNum())

	// recursion
	fmt.Println(factorial(3))

	// defer function starts when the enclosing function ends (main() in this case)
	// could be a clean up operation, close file, etc...
	//defer printTwo()
	//printOne()

	// safe division - catch expression
	fmt.Println("*** SAFE DIVISION")
	fmt.Println(safeDiv(3, 0))
	fmt.Println(safeDiv(3, 2))

	// throw/rise exception
	demPanic()
}

func demPanic() {

	defer func() {
		fmt.Println(recover())
	}()

	panic("this throws a PANIC message, throws error")
}

func safeDiv(num1, num2 int) int {

	defer func() {
		fmt.Println(recover())
	}()

	solution := num1 / num2
	return solution

}

func printOne() { fmt.Println(1) }
func printTwo() { fmt.Println(2) }

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func subtractThem(args ...int) int {
	finalValue := 0
	for _, value := range args {
		finalValue -= value
	}
	return finalValue
}

func next2values(number int) (int, int) {
	return number + 1, number + 2
}

func addAllNumsFromArray(numbers []float64) float64 {
	sum := 0.0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
