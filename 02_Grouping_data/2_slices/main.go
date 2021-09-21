package main

import "fmt"

func main() {
	// Slices

	// var slice []int = []int{1, 2, 3, 4, 5}
	// slice := []int{1, 2, 3, 4, 5} // initialization with composite literal
	slice := make([]int, 5, 10)

	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[3] = 4
	slice[4] = 5
	for i, x := range slice {
		fmt.Println("Element at index", i, "is", x)
	}
	fmt.Println("############################")

	// length and capacity of slice
	fmt.Println("Size:", len(slice))
	fmt.Println("Capacity:", cap(slice))
	fmt.Println("############################")

	// Delete an element from slice
	i := 2
	slice = append(slice[:i], slice[i+1:]...)
	fmt.Println("New slice after deleting element at index", i, " :", slice)
	fmt.Println("############################")

	// Append to a slice
	y := []int{11, 12, 13, 14, 15}
	new_slice := append(slice, y...)
	fmt.Println("New slice after appending:", new_slice)
	fmt.Println("############################")
}
