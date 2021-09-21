package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	time.Sleep(time.Second * 5)
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	arr := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(arr[:], c)

	fmt.Println("Point 1 executed...")
	s := <-c // receive from c
	fmt.Println("Point 2 executed...")

	fmt.Println(s)
}
