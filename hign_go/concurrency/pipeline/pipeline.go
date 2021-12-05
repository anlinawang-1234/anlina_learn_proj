package main

import (
	"fmt"
)

func main() {
	numChan := gen(1, 2, 3, 4, 5, 6, 7, 8, 9)
	for res := range sq(numChan) {
		fmt.Println(res)
	}
}

func gen(n ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, nums := range n {
			out <- nums
		}
		close(out)
	}()
	return out
}

func sq(n <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for nums := range n {
			out <- nums * nums
		}
		close(out)
	}()
	return out
}
