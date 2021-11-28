package main

import "fmt"

func main() {
	a := 10
	a = 100

	go func() {
		fmt.Println("a", a)
	}()
}
