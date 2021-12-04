package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 3)
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func(i int) {
			fun(i)
			time.Sleep(3 * time.Second)
			<-ch
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func fun(i int) {
	fmt.Println(i, "hello, world")
}
