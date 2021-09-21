package main

import "fmt"

func main() {
	c := make(chan int, 2)

	c <- 5
	c <- 6
	fmt.Println(<-c)
	fmt.Println(<-c)
}
