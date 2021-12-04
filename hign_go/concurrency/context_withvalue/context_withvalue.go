package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(), "name", "value")
	fmt.Println(ctx.Value("name"))
	f1(&ctx)
	f2(&ctx)
	timeoutCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	select {
	case <-timeoutCtx.Done():
		fmt.Println("时间到了....")
		fmt.Println(ctx.Value("f1"))
		fmt.Println(ctx.Value("f2"))
	}
}

func f1(ctx *context.Context) {
	fmt.Println("f1", (*ctx).Value("name"))
	*ctx = context.WithValue(*ctx, "f1", "函数1")
	return
}

func f2(ctx *context.Context) {
	fmt.Println("f2", (*ctx).Value("f1"))
	*ctx = context.WithValue(*ctx, "f2", "函数2")
}
