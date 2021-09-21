package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			ch1 <- "0.5 seconds"
		}
	}()
	go func() {
		for {
			time.Sleep(time.Millisecond * 2000)
			ch2 <- "2 seconds"
		}
	}()

	for {
		// fmt.Println(<-ch1)
		// fmt.Println(<-ch2)
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}

}
