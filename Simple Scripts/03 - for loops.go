package main

import "fmt"

func main() {
	// create iterator
	i := 1

	for i <= 10 {

		fmt.Println(i)
		i++
		// i = i + 1

	}

	// inline iterator declaration
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}
}
