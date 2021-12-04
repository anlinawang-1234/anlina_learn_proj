package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var Counter = 0

func main() {
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go count()
	}
	wg.Wait()
	fmt.Println("Counter", Counter)
}

func count() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		atomic.AddInt32(&int32(Counter), 1)
		time.Sleep(1 * time.Millisecond)
		//Counter++
	}
}
