package main

import "fmt"

func main() {
	// Array
	// var arr [5]int = [5]int{1, 2, 3, 4, 5}
	arr := [5]int{1, 2, 3, 4, 5}

	// Accessing an element
	fmt.Println(arr[2])
	fmt.Println("############################")

	// Traversing through array
	for i, x := range arr {
		fmt.Println("Element at index", i, "is", x)
	}
	fmt.Println("############################")

	// Modifying array
	arr[2] = arr[2] * 10
	fmt.Println(arr)
	fmt.Println("############################")

	// Slicing array
	fmt.Println(arr[1:3])
}
