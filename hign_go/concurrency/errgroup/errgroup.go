package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("g1")
		return nil
	})

	g.Go(func() error {
		fmt.Println("g2")
		//return errors.New("err")
		return nil
	})

	g.Go(func() error {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("g3")
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println("错误", err)
	}
	fmt.Println("ctx.Err", ctx.Err())
}
