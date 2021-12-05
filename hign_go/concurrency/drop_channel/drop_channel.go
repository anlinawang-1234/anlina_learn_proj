// 谁快就用谁的数据
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	//arrStr := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	arrInt := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range arrInt {
		value := v
		go func(a int) {
			select {
			case ch <- doWork(a):
				fmt.Println("写进去数据", value)
			default:
				fmt.Println("走默认", value)
			}
		}(value)
	}
	fmt.Println("ch的数据是", <-ch)
	time.Sleep(5 * time.Second)
}

func doWork(a int) int {
	//fmt.Println("执行了", str)
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)
	time.Sleep(time.Duration(num) * time.Millisecond)
	return a
}
