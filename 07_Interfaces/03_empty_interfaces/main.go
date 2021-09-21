package main

import (
	"fmt"
)

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "Hello World"
	i := 55
	strt := struct {
		name string
	}{
		name: "David",
	}

	describe(i)
	describe(s)
	describe(strt)
}
