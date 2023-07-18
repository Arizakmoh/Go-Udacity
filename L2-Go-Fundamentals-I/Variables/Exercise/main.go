package main

import "fmt"

/*Then, within the main() function:

Use the var keyword to declare a language variable that points to the string "Go"
Use the const keyword to declare a released variable that points to the integer 2009
Use shorthand notation to declare a isAProgrammingLanguage variable that points to the boolean true
Print the values of the three variables (i.e., language, released, isAProgrammingLanguage) to the console and view the output!*/

func main() {
	var language string = "Go"
	const released int = 2009
	isAProgrammingLanguage := true

	fmt.Println("Language : ", language)
	fmt.Println("Released : ", released)
	fmt.Println("isAProgrammingLanguage : ", isAProgrammingLanguage)

}
