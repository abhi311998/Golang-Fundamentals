package main

import "fmt"

func main() {
	a := 120
	foo(a)
	fmt.Println("Value of a after calling foo:", a)

	fmt.Println()
	bar(&a)
	fmt.Println("Value of a after calling bar:", a)
}

func foo(a int) {
	a = a * 10
	fmt.Println("Value of a inside foo function:", a)
}

func bar(a *int) {
	*a = (*a) * 10
	fmt.Println("Value of a inside bar function:", *a)
}
