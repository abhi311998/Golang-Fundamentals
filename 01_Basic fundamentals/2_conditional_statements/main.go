package main

import (
	"fmt"
	"time"
)

func foo() {
	/*
		if  var declaration;  condition {
			// code to be executed if condition is true
		}
	*/
	if x := 100; x == 100 {
		fmt.Println("Other way of using if statement")
	}
}

func main() {
	// x := 100

	// if x == 50 {
	// 	fmt.Println("A")
	// } else if x == 100 {
	// 	fmt.Println("B")
	// } else {
	// 	fmt.Println("C")
	// }

	foo()
	switches()
}

func switches() {
	today := time.Now()

	switch today.Day() {
	case 1, 2, 3, 4, 5, 6, 7:
		fmt.Println("First week of the month.")
	case 8, 9, 10, 11, 12, 13, 14:
		fmt.Println("Second week of the month.")
	case 15, 16, 17, 18, 19, 20, 21:
		fmt.Println("Three week of the month.")
	case 22, 23, 24, 25, 26, 27, 28:
		fmt.Println("Fourth week of the month.")
	case 29, 30, 31:
		fmt.Println("Last week of the month..")
	default:
		fmt.Println("No information available for that day.")
	}
}
