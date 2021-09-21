package main

import "fmt"

func main() {
	e1 := struct {
		FirstName, LastName string
		Age                 int
	}{
		FirstName: "Jack",
		LastName:  "Snyder",
		Age:       32,
	}

	fmt.Println("struct:", e1)
	fmt.Println("Employee first name:", e1.FirstName)
	fmt.Println("Employee last name:", e1.LastName)
	fmt.Println("Employee age:", e1.Age)
}
