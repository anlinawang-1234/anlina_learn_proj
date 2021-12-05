package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	dataCh := make(chan int, 5)
	ctx, cancel := context.WithCancel(context.Background())

	// write
	wg.Add(3)
	go func() {
		defer wg.Done()
		i := 1
		for {
			select {
			case <-ctx.Done():
				fmt.Println("退出写...", ctx.Err())
				return
			default:
				time.Sleep(1 * time.Second)
				i++
				dataCh <- i
			}
		}
	}()

	// read
	go func() {
		defer wg.Done()
		for {
			select {
			case value := <-dataCh:
				fmt.Println("read value: ", value)
			case <-ctx.Done():
				fmt.Println("退出读...", ctx.Err())
				return
			}
		}
	}()

	// cancel
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 5)
		cancel()
	}()

	wg.Wait()
}
