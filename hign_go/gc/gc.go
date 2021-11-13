package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"

)

func main(){
	//gc1()
	//gcTrace()
	traceOut()
	fmt.Sprintf("hello ")
}

// 生成trace 图 使用 go tool trace trace.out
func traceOut(){
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	keepAlloc()
}

var cache = map[interface{}]interface{}{}
func keepAlloc(){
	for i := 0; i < 10000; i++{
		cache[i] = make([]byte, 1<<10)
	}
}

func gcTrace(){
	// 查看gc信息
	for i := 0; i < 10000; i++{
		_ = make([]byte, 1<<20)
	}
}
func gc1() {
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
