package main

import (
	. "fmt"
	"time"
)

var ch = make(chan int, 1)
var A = 0

func main() {
	demo1()
}

func demo1() {
	go add()
	//go read()
	<-ch
	Println(A)
}

func add() {
	A = 666
	time.Sleep(time.Second * 1)
	ch <- 1
}
