package main

import "fmt"

func main() {
	closure()
}

// 闭包
func do_closure() func() {
	return func() {
		a := 1
		a++ // a escapes to heap
		fmt.Println("a = ", a)
	}
}
func closure() {
	d1 := do_closure()
	d1()
	d1()
	d1()
	d1()
	d1()
}
