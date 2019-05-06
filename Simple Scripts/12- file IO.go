package main

import (
	"fmt"
	"os"
	//"log"
	//"io/ioutil"
	//"strconv"
	"log"
	"io/ioutil"
)

func main() {

	file, err := os.Create("samp.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("This is some random text")
	file.Close()

	stream, err := ioutil.ReadFile("samp.txt")

	readString := string(stream)
	fmt.Println(readString)

}
