package main

import "fmt"

func main() {
	foo()
	defer func_1()
	defer func_2()
	defer func_3()
	bar()
}

func foo() {
	fmt.Println("Inside foo")
}

func bar() {
	fmt.Println("Inside bar")
}

func func_1() {
	fmt.Println("Inside func_1")
}

func func_2() {
	fmt.Println("Inside func_2")
}

func func_3() {
	fmt.Println("Inside func_3")
}
