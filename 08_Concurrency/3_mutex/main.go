package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0

	const gs = 1000
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count:", counter)
}
