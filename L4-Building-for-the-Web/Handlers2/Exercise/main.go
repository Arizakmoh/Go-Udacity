/*
Exercise: Handlers II
In this exercise, you'll gain even more practice building an HTTP server.

First, navigate to / L4. Building for the Web / 2. Handlers II / Exercise / main.go

To begin:

1: Uncomment the commented code.
2 : Create a single handler index() that:
	A: Sets the content type to HTML
	B: Iterates through the cityPopulations map and displays an H2 header of the city name and its population. Notes: each key-value pair should be displayed on separates lines.
	C: Returns the appropriate HTTP status code upon success
3: Register your handler.
4: Serve your program on port 3000.
5: Test out your endpoint via cURL or Postman. Do you see the expected output?
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

var cityPopulations = map[string]uint32{
	"Tokyo":       37435191,
	"Delhi":       29399141,
	"Shanghai":    26317104,
	"Sao Paulo":   21846507,
	"Mexico City": 21671908,
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset-utf-8")
	w.WriteHeader(http.StatusOK)

	for m, city := range cityPopulations {
		fmt.Fprintf(w, "<h2>%s :  %d </h2>", m, city)
	}
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server Started at : ", time.Now())

	http.ListenAndServe(":3000", nil)
}
