package main

import (
	"fmt"
	"net/http"
	"sync"
)

func newBuf()[]byte{
	return make([]byte, 10<<20)
}

func main(){
	newBuffer := sync.Pool{

	}
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	http.HandleFunc("/example", func(writer http.ResponseWriter, request *http.Request) {
		b := newBuf()

		for idx := range b{
			b[idx] = 1
		}
		fmt.Fprintf(writer, "done +%v", request.URL.Path)
	})
	http.ListenAndServe(":8080", nil)
}

