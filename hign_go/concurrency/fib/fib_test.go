package main

import (
	"fmt"
	"testing"
)

func main() {
	b := testing.B{}
	BenchmarkFib(&b)
}

func BenchmarkFib(b *testing.B) {
	i := 0
	for ; i < b.N; i++ {
		//fmt.Println("i=", i, b.N)
		fib(30)
	}
	fmt.Println("i", i, b.N)
}
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}
