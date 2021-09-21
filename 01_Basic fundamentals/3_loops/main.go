package main

import "fmt"

func main() {

	// For loop
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// While loop
	i := 1
	for i <= 10 {
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	// Infinite loop
	i = 1
	for {
		fmt.Print(i, " ")
		if i == 10 {
			break
		}
		i++
	}
	fmt.Println()

	// range keyword
	arr := []int{1, 2, 3, 4, 5}
	for index, val := range arr {
		fmt.Println(index, val)
	}
}
