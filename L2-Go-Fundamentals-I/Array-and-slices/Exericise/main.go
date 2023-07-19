/*Exercise: Arrays and Slices
Exercise: Arrays and Slices
A slice is an excellent data structure that stores a dynamically-sized amount of elements. In this exercise, you'll practice creating a basic slice, use your new skills to access and modify the slice.

First, navigate to /L2. Go Fundamentals I / 4. Arrays and Slices / Exercise / main.go

Then, within the main() function:

Create a slice of strings called languages with the following four elements: "Go", "JavaScript", "Ruby", and "Python"
Print the entire slice to the console
Print the length of the slice to the console
Print the first element of the slice
Print a slice of just the second and third elements (hint: you'll need to provide low-bound and high-bound indices!)
Append a new element to the language slice: "PHP"
Print the newly-modified slice to the console
If you get stuck, feel free to review the provided Solution as well.*/

package main

import "fmt"

func main() {
	var languages = []string{"Go", "JavaScript", "Ruby", "Python"}
	fmt.Println(languages)
	fmt.Println(len(languages))
	fmt.Println(languages[0])
	fmt.Println(languages[1:3])
	languages = append(languages, "PHP")
	fmt.Println(languages)
}
