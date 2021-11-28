package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	tr := NewTrack()
	go tr.Run()
	ctx, cancel := context.WithDeadline(context.Background(), time.Now())
	defer cancel()
	tr.Event(context.Background(), "test1")
	tr.Event(context.Background(), "test2")
	tr.Event(context.Background(), "test3")
	tr.Shutdown(ctx)
}

func NewTrack() *Track {
	return &Track{
		DataCh: make(chan string, 10),
		Status: make(chan struct{}),
	}
}

type Track struct {
	DataCh chan string
	Status chan struct{}
}

func (t *Track) Run() {
	for data := range t.DataCh {
		//time.Sleep(time.Millisecond*1)
		fmt.Println(data)
	}
	t.Status <- struct{}{}
}

func (t *Track) Event(ctx context.Context, data string) error {
	select {
	case t.DataCh <- data:
		fmt.Println("写入数", data)
		return nil
	case <-ctx.Done():
		fmt.Println("被取消了")
		return errors.New("canceled")
	}
	return nil
}

func (t *Track) Shutdown(ctx context.Context) {
	fmt.Println("开始shutdown")
	close(t.DataCh)
	select {
	case <-t.Status:
		fmt.Println("收到退出信号")
	case <-ctx.Done():
		fmt.Println("shutdown被取消了")
	}
}
