package main

import (
	"context"
	"fmt"
	"time"
)

type result struct {
	Record string
	Err    error
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	ch := make(chan result)

	go func() {
		record, err := search("xxx")
		ch <- result{Record: record, Err: err}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("处理超时了")
	case result := <-ch:
		fmt.Println("获取到了结果", result)
	}
}

func search(tmp string) (string, error) {
	time.Sleep(time.Millisecond * 100)
	return "some search", nil
}
