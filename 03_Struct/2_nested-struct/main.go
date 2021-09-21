package main

import "fmt"

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              Salary
}

// Nested struct

func main() {
	e1 := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: Salary{
			Basic: 15000.00,
			HRA:   5000.00,
			TA:    2000.00,
		},
	}
	fmt.Println("Nested struct:", e1)
}
