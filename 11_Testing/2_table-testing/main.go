package main

import "fmt"

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("2 + 3 =", mySum(2, 3))
	fmt.Println("2 + 3 + 4 =", mySum(2, 3, 4))
	fmt.Println("200 + 33 =", mySum(200, 33))
}
