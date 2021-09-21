package main

import "fmt"

type Employee struct {
	FirstName, LastName string
	Salary              float32
}

func main() {
	// e1 := Employee{
	// 	FirstName: "Jack",
	// 	LastName:  "Snyder",
	// 	Salary:    32000,
	// }
	// fmt.Println("struct:", e1)
	// e1.printEmployee()
	// incrementSalary(&e1)
	// e1.printEmployee()

	// Anonymous Function
	t := func() string {
		return "Anonymous Funtion..."
	}()
	fmt.Println(t)

	// Returning a func from a func
	t2 := dummy()
	fmt.Println(t2())
}

// // Functions
// func incrementSalary(e *Employee) {
// 	e.Salary = e.Salary * (1.1)
// }

// // Methods
// func (e Employee) printEmployee() {
// 	fmt.Println("Employee first name:", e.FirstName)
// 	fmt.Println("Employee last name:", e.LastName)
// 	fmt.Println("Employee Salary:", e.Salary)
// 	fmt.Println()
// }

func dummy() func() string {
	return func() string {
		return "Returning a func from a func..."
	}
}
