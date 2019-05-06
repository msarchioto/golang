package main

import "fmt"
import "net/http"

func main() {

	http.HandleFunc("/", handler)

	http.HandleFunc("/earth", handler2)

	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello World\n")
}

func handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello Earth\n")
}