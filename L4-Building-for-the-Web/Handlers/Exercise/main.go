/*Exercise: Handlers I
In this exercise, you'll leverage the net/http package to support two endpoints: one that displays a simple message on the page, and other than displays a list of the five most populous cities in the world.

First, navigate to /L4. Building for the Web / 1. Handlers I / Exercise / main.go

To begin:

1 : Uncomment the commented code.
2 : Create two handler functions:
	A : index(), which responds with a simple text greeting (of your choice) on the page
	B : cityList(), which responds with "List of most populous cities." Then, below, lists each of the cities in the cities slice on separate lines.
3 : Register your handlers. Hint: Which method in the net/http should you use?
4 : Serve your program on port 3000.
5 :s Test out your endpoints via cURL or Postman. Do you see the expected output?*/

package main

import (
	"fmt"
	"net/http"
)

var cities = []string{"Tokyo", "Delhi", "Shanghai", "Sao Paulo", "Mexico City"}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index Page")
}

func cityList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All the Cities")

	for _, city := range cities {
		fmt.Fprintf(w, "%s\n", city)
	}
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/cityList", cityList)

	fmt.Println("Server Started")
	http.ListenAndServe(":3000", nil)
}
