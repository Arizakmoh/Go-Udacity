package main

import "fmt"

func getRectangleArea(width int, height int) string {
	area := width * height

	if area < 50 {
		return fmt.Sprintf("he area is %d, which is less than 50", area)
	} else {
		return fmt.Sprintf("The area is %d, which is greater than or equal to 50", area)
	}
}

func main() {
	result := getRectangleArea(5, 7)
	fmt.Println(result)
}
