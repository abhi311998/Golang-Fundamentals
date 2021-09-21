package main

import (
	"fmt"
	"time"
)

func main() {
	foo()
	go bar()

	// time.Sleep(time.Second)
	fmt.Println("Inside function - main")
}

func foo() {
	time.Sleep(time.Second)
	fmt.Println("Inside function - foo")
}

func bar() {
	fmt.Println("Inside function - bar")
}
