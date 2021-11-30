package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//incorrect1()
	incorrect2()
}

var a, b int

func f1() {
	a = 1
	b = 2
}
func g1() {
	fmt.Println(b)
	fmt.Println(a)
}
func incorrect1() {
	go f1()
	g1()
}

var str string
var done bool
var once sync.Once

func setup() {
	str = "hello, world"
	done = true
}
func doPrint() {
	once.Do(func() {
		fmt.Println(str)
	})
}
func incorrect2() {
	setup()
	go doPrint()
	go doPrint()
	time.Sleep(time.Second * 3)
}
