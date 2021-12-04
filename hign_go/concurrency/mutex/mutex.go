package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan int, 1)
	var mu sync.Mutex
	var g1, g2 int

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				mu.Lock()
				time.Sleep(100 * time.Millisecond)
				g1++
				mu.Unlock()
			}
		}
	}()

	for i := 0; i < 10; i++ {
		mu.Lock()
		g2++
		mu.Unlock()
		time.Sleep(100 * time.Millisecond)
	}
	done <- 1
	fmt.Println("g1=", g1, "g2=", g2)
}
