/*Exercise: Ranges
Exercise: Ranges
A range is a great way to iterate through a data structure such a slice. In this exercise, you'll create a reduce() function that returns the sum of all integers in a slice (i.e., adds all values of the slice together).

First, navigate to /L2. Go Fundamentals I / 6. Ranges / Exercise / main.go

To begin:

Create a reduce() function that takes in a single argument: a slice of integers. The function itself should return an integer (i.e., the result of adding all integers in the slice together).
Use a range to iterate through the slice and adds all numbers together. Hint: Create a sum variable that cumulatively builds as you iterate through the slice.
Uncomment the commented lines and run your code! Does 54 get outputted to the console?
While the amount code for this exercise may not be too large, we recommend taking your time to develop the logic for it. If you get stuck, you are welcome to review the solution provided.*/

package main

import "fmt"

func main() {
	number := []int{10, 50, 13}

	result := reduce(number)
	fmt.Println(result)
}

func reduce(number []int) int {
	var sum int

	for _, num := range number {
		sum = sum + num
	}

	return sum
}
