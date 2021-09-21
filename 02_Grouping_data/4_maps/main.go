package main

import "fmt"

func main() {
	mp := make(map[string]int)
	// mp := map[string]int{}
	mp["one"] = 1
	mp["two"] = 2
	mp["three"] = 3

	// Accessing value using key
	fmt.Println(mp["three"])
	fmt.Println("############################")

	// Traversing through map
	for key, val := range mp {
		fmt.Println("Key:", key, ", Value:", val)
	}
	fmt.Println("############################")

	// Delete a pair
	fmt.Println("Before deleting:", mp)
	delete(mp, "two")
	fmt.Println("After deleting:", mp)
}
