package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	emp1 := Employee{
		"John",
		"Doe",
		1,
	}

	emp2 := Employee{
		firstName: "Jake",
		lastName:  "Deer",
		id:        2,
	}

	var emp3 Employee
	emp3.firstName = "Dude"
	emp3.lastName = "Just Dude"
	emp3.id = 3

	fmt.Println(emp1, emp2, emp3)
}
