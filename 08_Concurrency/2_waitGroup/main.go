package main

import (
	"fmt"
	"sync"
)

// https://play.golang.org/p/DWGO5aK2qwf

func main() {
	counter := 0

	const gs = 1000
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count:", counter)
}
