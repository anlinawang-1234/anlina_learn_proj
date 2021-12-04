package main

import (
	"fmt"
	"sync"
)

type DeadLock struct {
	Value int32
	Mutex sync.Mutex
}

var wg sync.WaitGroup

func main() {
	//dl := DeadLock{Value: 1}
	mapDl := make(map[int32]*DeadLock)
	mapDl[1] = &DeadLock{Value: 1}
	mapDl[2] = &DeadLock{Value: 2}
	mapDl[3] = &DeadLock{Value: 3}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go updateValue(&mapDl)
	}
	wg.Wait()

	for _, v := range mapDl {
		fmt.Println(v.Value)
	}
}

func updateValue(mapDl *map[int32]*DeadLock) {
	defer wg.Done()

	for _, v := range *mapDl {
		v.Mutex.Lock()
		v.Value++
		v.Mutex.Unlock()
	}
}
