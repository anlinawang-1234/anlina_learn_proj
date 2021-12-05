package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan string)

	go func() {
		defer func() {
			wg.Done()
			fmt.Println("write out...")
		}()
		ch <- "foo"
	}()

	go func() {
		defer func() {
			wg.Done()
			fmt.Println("read out...")
		}()
		time.Sleep(time.Second * 1)
		fmt.Println("read value :", <-ch)
	}()

	wg.Wait()
}
