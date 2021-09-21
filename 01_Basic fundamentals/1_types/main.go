package main

import "fmt"

func main() {
	// Primitive data types
	var bo bool = true
	fmt.Printf("Type: %T, Value: %v\n", bo, bo)

	var b byte = 8
	fmt.Printf("Type: %T, Value: %v\n", b, b)

	var i int = 54
	fmt.Printf("Type: %T, Value: %v\n", i, i)

	var f float32 = 54.678
	fmt.Printf("Type: %T, Value: %v\n", f, f)

	var s string = "abcdef"
	fmt.Printf("Type: %T, Value: %v\n", s, s)

	//////////////////////////////////////////////////////////
	// User defined types
	type name string // underlying type of name is string, so we can use it to initialize any var of type string
	var n name = "asdfg"
	fmt.Printf("Type: %T, Value: %v\n", n, n)

}
