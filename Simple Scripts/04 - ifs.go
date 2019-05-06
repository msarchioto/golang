package main

import "fmt"

func main() {

	yourAge := 18

	if yourAge >= 16 {
		fmt.Println("You can drive")
	} else if yourAge >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can't have fun")
	}

	switch yourAge {

	case 16:
		fmt.Println("Go Drive")
	case 18:
		fmt.Println("Go Vote")
	default:
		fmt.Println("Go Have Fun")

	}
}
