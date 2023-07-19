/*Exercise: Goroutines
Exercise: Goroutines
By leveraging Goroutines, we can create functions that run independently of each other. This way, we can have code run concurrently -- leading to faster and more performant programs.

In this exercise, you'll prove this statement by using the time package and running two function invocations concurrently.

First, navigate to /L2. Go Fundamentals II / 4. Goroutines / Exercise / main.go

First, uncomment the starter code. Then, to begin:

Import the time package.
1: Create a function that takes in a slice of strings, then simply iterates through that slice and prints out each string in it. Hint: Which slice(s) should you pass into this function?
2: Using the go keyword, invoke the function concurrently to print out the values of both slices.
3: Using the time package, keep track of the start time as well as the total duration of execution. Review the package's documentation (above) to determine how to get the current time as well as the elapsed time since that start time. Hint: When should you start the timer? When should you end it?
4: Run your main.go program with and without the use of Goroutines. Which is faster?*/

package main

import (
	"fmt"
	"time"
)

func iteratesAndPrint(slice []string) {
	for _, item := range slice {
		fmt.Println(item)
	}
}

func main() {

	// Without Gorouting

	startTime := time.Now()

	slice1 := []string{"Mohamed", "Abdirizak"}
	slice2 := []string{"Maryama", "Nadiifo"}

	fmt.Println("Calling or Printing Without Gorouting")
	iteratesAndPrint(slice1)
	iteratesAndPrint(slice2)

	duration := time.Since(startTime)
	fmt.Printf("Duration Without a Gorouting %d\n", duration)

	// With Gorouting

	// Without Gorouting
	fmt.Println("***********************")
	startTime = time.Now()

	fmt.Println("Calling or Printing Without Gorouting")
	go iteratesAndPrint(slice1)
	go iteratesAndPrint(slice2)
	time.Sleep(time.Millisecond) // Wait for Goroutines to finish

	duration = time.Since(startTime)
	fmt.Printf("Duration With a Gorouting %d\n", duration)

}
