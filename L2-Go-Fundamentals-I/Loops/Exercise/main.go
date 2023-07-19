/* A loop is a great way to repeat blocks of code until a condition is met. In this exercise, you'll implement the classic software engineer interview question: Fizzbuzz. That is, given an integer n, you'll print a slice of strings of 1 through n where:

If the integer is divisble by 3, the value will be "Fizz"
If the integer is divisble by 5, the value will be "Fizz"
If the integer is divisible by both 3 and 5, the value will be "Fizzbuzz"
If none of the above is true, the value will be the integer itself
For example: Given an n of 15, the result would be:

["1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"]
First, navigate to /L2. Go Fundamentals I / 5. Loops / Exercise / main.go

To begin:

Create a fizzbuzz() function that takes in an integer n and returns a slice of strings.
Write a for loop that iterates from 1 through n (inclusive) and satisfies the list of requirements above. Hint: How can you conditionally check that an integer in the loop is both divisible by 3 and 5? When should you make that check?
Return a slice of strings containing all the outputs. Hint: Be sure to convert any integers into strings as needed!
In the main() function, print the output of invoking your new fizzbuzz() function! Feel free to pass any value of n into your function invocation.
This exercise is a bit more complex that the previous exercises in this lesson. We recommend you taking your time building the logic for this one! If you need some inspiration, however, feel free to review the solution.*/

/*package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzbuzz(15))
}
func fizzbuzz(n int) []string {
	var result []string

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "Fizzbuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzbuzz(15))
}
func fizzbuzz(n int) []string {
	var result []string

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "Fizzbuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(n))
		}
	}

	return result

}
