package main

import (
	"context"
	_ "fmt"
	"time"
)

var ch = make(chan int, 1)
var A = 0

func main() {
	//callf()
	limitChannel()
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*5000)
	defer cancel()
	select {
	case <-ctx.Done():
		return
	}
}

// A receive from an unbuffered channel happens before the send on that channel completes.
var ch1 = make(chan int, 1)
var a string

func f1() {
	a = "hello, world"
	close(ch1)
}
func f2() {
	<-ch1
	println(a)
}
func callf() {
	go f1()
	go f2()
	time.Sleep(time.Second * 1)
}

// This program starts a goroutine for every entry in the work list,
//but the goroutines coordinate using the limit channel to ensure that at most
//three are running work functions at a time.
var ch3 = make(chan int, 3)

func limitChannel() {
	for i := 0; i < 10; i++ {
		index := i
		go func() {
			ch3 <- 1
			println(index, "号执行 hello, world")
			time.Sleep(time.Second * 1)
			<-ch3
		}()
	}
}
