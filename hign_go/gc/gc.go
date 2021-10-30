package main

import (
	"fmt"
	"runtime"
	"time"

)

func main(){
	startTime := time.Now()
	i := 0
	go func() {
		for true {
			i++
		}
	}()

	time.Sleep(time.Millisecond)
	runtime.GC()
	fmt.Println("OK", time.Now().Sub(startTime))
}
