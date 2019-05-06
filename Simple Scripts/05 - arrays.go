package main

import "fmt"

func main() {

	var favNums2 [5]float64

	favNums2[0] = 163
	favNums2[1] = 7387
	favNums2[2] = 23
	favNums2[3] = 3.141
	favNums2[4] = .1234

	fmt.Println(favNums2[3])

	favNums3 := [5]float64{1, 2, 3, 4, 5}

	for i, value := range favNums3 {
		fmt.Println(value, i)
	}

	// undefined index
	for _, value := range favNums3 {
		fmt.Println(value)
	}

	// slices (no size)
	numSlice := []int{5, 4, 3, 2, 1}

	numSlice2 := numSlice[3:5]

	fmt.Println("numslice2[0] =", numSlice2[0])
	fmt.Println(numSlice2)

	fmt.Println("numslice2[0:2] =", numSlice[:2])
	fmt.Println("numslice2[2:] =", numSlice[2:])

	// slice with no defined values (type, default 0 for the first 5 elements, max values)
	numSlice3 := make([]int, 5, 10)
	fmt.Println(numSlice3)

	copy(numSlice3, numSlice)
	fmt.Println(numSlice3)

	numSlice3 = append(numSlice3, 0, -1)
	fmt.Println(numSlice3)

}
