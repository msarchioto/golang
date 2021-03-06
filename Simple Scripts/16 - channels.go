package main

import "fmt"
import "time"
import (
	"strconv"
)

// vars

var pizzaNum = 0
var pizzaName = ""

func main() {

	stringChan := make(chan string)

	for i := 0; i < 3; i++ {
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)

		time.Sleep(time.Millisecond*5000)
	}


}

func makeDough(stringChan chan string){
	pizzaNum++
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)
	fmt.Println("Make Dough and send for sauce")

	stringChan <- pizzaName
	time.Sleep(time.Millisecond*10)
}

func addSauce(stringChan chan string){
	pizza := <- stringChan
	fmt.Println("Add sauce and send", pizza, "for topping")

	stringChan <- pizzaName
	time.Sleep(time.Millisecond*10)
}

func addToppings(stringChan chan string){
	pizza := <- stringChan
	fmt.Println("Add toppings to", pizza, "and ship")

	time.Sleep(time.Millisecond*10)
}