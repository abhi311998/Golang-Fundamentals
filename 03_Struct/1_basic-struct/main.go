package main

import "fmt"

type Employee struct {
	FirstName, LastName string
	Age                 int
}

func main() {
	e1 := Employee{
		FirstName: "Jack",
		LastName:  "Snyder",
		Age:       32,
	}
	fmt.Println("struct:", e1)
	e1.printEmployee()
}

// Methods
func (e Employee) printEmployee() {
	fmt.Println("Employee first name:", e.FirstName)
	fmt.Println("Employee last name:", e.LastName)
	fmt.Println("Employee age:", e.Age)
}
