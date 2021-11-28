package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	//badCase()
	//selfControl()
	goodCase()
}

var done = make(chan error, 2)
var stop = make(chan struct{})

func goodCase() {
	go func() {
		done <- serverApp(stop)
	}()
	go func() {
		done <- serverMonitor(stop)
	}()

	stopped := false
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Println("错误了", err)
			if !stopped {
				stopped = true
				close(stop)
			}
		}
	}
}

var addr1 = ":8080"

func serverApp(stop chan struct{}) error {
	defer func() {
		fmt.Println("8080 要退出了")
	}()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello, 我来执行了", r)
		int := rand.Intn(10)
		fmt.Println("int", int)
		if int == 6 {
			done <- errors.New("错误了")
		}
	})
	http.HandleFunc("/", handler)
	server := &http.Server{
		Addr:    addr1,
		Handler: handler,
	}

	go func() {
		<-stop
		fmt.Println("8080 我收到了退出要求了")
		server.Shutdown(context.Background()) // todo 这里没有调试好，这里收到stop信号只有没有退出
	}()

	return http.ListenAndServe(":8080", nil)
}

var addr2 = "8081"

func serverMonitor(stop chan struct{}) error {
	defer func() {
		fmt.Println("8081 要退出了")
	}()

	server := &http.Server{
		Addr: addr2,
	}
	go func() {
		<-stop
		fmt.Println("8081 我收到了退出要求了")
		server.Shutdown(context.Background())
	}()
	return http.ListenAndServe(":8081", nil)
}

// ListenAndServe是阻塞函数，自己进行线程控制
func selfControl() {
	chan1 := make(chan int, 0)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello, http func", w, "请求参数", r.Form)
		chan1 <- 8888
	})

	go func() {
		select {
		case num := <-chan1:
			fmt.Println("读取到数据", num)
		}
	}()

	// 监听端口
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("监听失败")
	}
}

func badCase() {
	chan1 := make(chan int, 1)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello, http func ", w)
	})

	go func() {
		err := http.ListenAndServe(":8080", nil)
		fmt.Println("监听端口err", err)
		if err == nil {
			fmt.Println("8080端口")
			chan1 <- 1000
			fmt.Println("写入数据")
		}
	}()

	select {
	case num := <-chan1:
		fmt.Println("读取到数", num)
	}
}
