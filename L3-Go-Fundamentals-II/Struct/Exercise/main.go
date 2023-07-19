/*Exercise: Structs
Exercise: Structs
In this exercise, you'll gain practice creating two different structs. Then, you'll initialize two different structs using two different strategies: by creating a struct literal, and by using the new keyword.

First, navigate to /L2. Go Fundamentals II / 2. Structs / Exercise / main.go

To begin:

1 : Create a Student struct with two fields: id and name.
2: Create a Classroom struct with four fields: id, capacity, subject, and studentList. Hint: If you were to use a slice for the studentList, what type of data should be in that slice?
3: Within the main() function, declare a c1 variable and create a Classroom struct literal with any values that you choose (i.e., for the fields defined in the previous steps). Be sure to initialize at least two students for the studentList.
4: Still within the main() function, declare a c2 variable whose value is the result of initializing another Classroom -- this time using the new keyword. Again, you may assign any values you choose to the fields defined in steps 1-2.
5: Print c1 and c2 and view your output!
*/

package main

import "fmt"

// 1 : Create a Student struct with two fields: id and name.
type student struct {
	id   uint16
	name string
}

//2: Create a Classroom struct with four fields: id, capacity, subject, and studentList. Hint: If you were to use a slice for the studentList, what type of data should be in that slice?

type Classroom struct {
	id          uint16
	capacity    uint16
	subject     string
	studentList []student
}

func main() {
	//3: Within the main() function, declare a c1 variable and create a Classroom struct literal with any values that you choose (i.e., for the fields defined in the previous steps). Be sure to initialize at least two students for the studentList.

	c1 := Classroom{
		id:       10,
		capacity: 10,
		subject:  "English",
		studentList: []student{
			{
				id:   1,
				name: "ABDIRIZAK",
			},
		},
	}

	//4: Still within the main() function, declare a c2 variable whose value is the result of initializing another Classroom -- this time using the new keyword. Again, you may assign any values you choose to the fields defined in steps 1-2.

	C2 := new(Classroom)
	C2.id = 11
	C2.capacity = 12
	C2.subject = "Somalia"
	C2.studentList = []student{
		{
			id:   2,
			name: "warsame",
		},
	}

	//5: Print c1 and c2 and view your output!

	fmt.Println(c1)
	fmt.Println(C2)

}
