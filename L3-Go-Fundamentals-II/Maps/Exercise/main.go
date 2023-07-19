/*Exercise: Maps
Maps are versatile and powerful data structures, with fast lookups and values that can be fetched, updated, or deleted via their keys. In this exercise, you'll create a map that represents the available courses at a local university. You'll then iterate through the map with a range and print out the courses matching a specific criterium.

First, navigate to /L2. Go Fundamentals II / 1. Maps / Exercise / main.go

Then, within the main() function:

1 : Create a map called courses containing seven string values with sequential integer keys (starting at 1). The string value should be the name of the course, while the integer key will represent the course ID. The courses are: Calculus, Biology, Chemistry, Computer Science, Communications, English, and Cantonese.
2 : Using a range, iterate over the courses map and print out only the courses that begin with the letter "C" (hint: this method should help!)
3 : Update the fourth key-value pair in this map to have a different value. That is, instead of "Computer Science," reassign the value to "Algorithms" without directly modifying the code you had written in Step 1 above
4: Using similar notation, add an eighth key-value pair to the map: "Spanish"
5 : Next, delete the first key-value pair of the map (again, without directly modifying the code you had written in Step 1)
6 : Print the new map to the console. Do you see the following output?*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	//1 : Create a map called courses containing seven string values with sequential integer keys (starting at 1). The string value should be the name of the course, while the integer key will represent the course ID. The courses are: Calculus, Biology, Chemistry, Computer Science, Communications, English, and Cantonese.

	courses := map[uint16]string{
		1: "Calculus",
		2: "Biology",
		3: "Chemistry",
		4: "Computer Science",
		5: "Communications",
		6: "English",
		7: "Cantonese",
	}

	//2 : Using a range, iterate over the courses map and print out only the courses that begin with the letter "C" (hint: this method should help!)

	for id, course := range courses {
		if strings.HasPrefix(course, "C") {
			fmt.Println(id, course)
		}
	}
	//3 : Update the fourth key-value pair in this map to have a different value. That is, instead of "Computer Science," reassign the value to "Algorithms" without directly modifying the code you had written in Step 1 above
	courses[4] = "Algorithms"

	//4: Using similar notation, add an eighth key-value pair to the map: "Spanish"
	courses[8] = "Spanish"

	//5 : Next, delete the first key-value pair of the map (again, without directly modifying the code you had written in Step 1)
	delete(courses, 1)

	//	6 : Print the new map to the console. Do you see the following output?*/

	fmt.Println(courses)

}
