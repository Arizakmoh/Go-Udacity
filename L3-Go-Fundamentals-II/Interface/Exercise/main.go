/*Exercise: Interfaces
Exercise: Interfaces
The main objective in this exercise is to re-write the provided starter code in such a way that uses an interface to produce the same output.

First, navigate to /L2. Go Fundamentals II / 3. Interfaces / Exercise / main.go

Then, before you begin, view the starter code in the main.go file and see if you can determine what the code is trying to accomplish. What is being calculated here? How is it implemented? Be sure to run the program and see what gets outputted to the console.

When you're ready to begin:

1: Create a basic interface for geometric shapes. This interface will be implemented on the Rectangle and Square types
2 : Implement the method(s) listed in the interface
3: Create a generic method into which we can pass in any initialized shape (of any of the two types) and have it return the perimeter of that shape
4 : Modify the print statements in the main() function to print out the perimeters of both initialized shapes!*/

package main

import (
	"fmt"
)

// Step 1 - Define the interface 
type Shape interface {
	Perimeter() float64
}

// Step 2 - Implement the interface 
type Rectangle struct {
	length, width float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

// Step 3 - Implement the interface
type Square struct {
	side float64
}

func (s Square) Perimeter() float64 {
	return 4 * s.side
}

// Step 4 - Create a generic function to calculate the perimeter 
func CalculatePerimeter(shape Shape) float64 {
	return shape.Perimeter()
}

func main() {
	rect := Rectangle{length: 10, width: 20}
	square := Square{side: 4}

	// Step 5 - Print the perimeters of both shapes
	fmt.Println("Rectangle Perimeter:", CalculatePerimeter(rect))
	fmt.Println("Square Perimeter:", CalculatePerimeter(square))
}