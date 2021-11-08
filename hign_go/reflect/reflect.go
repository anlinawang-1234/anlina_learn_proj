package main

import (
	"fmt"
	"reflect"
)

type Execute struct {
}

func main(){
	fmt.Println("hello")
	myExec := Execute{}
	execType := reflect.TypeOf(&myExec)
	for i := 0; i < execType.NumMethod(); i++{
		fmt.Println("method", i+1, execType.Method(i))
		reflect.ValueOf(&myExec).MethodByName(execType.Method(i).Name).Call(nil)
	}
	reflect.ValueOf(&myExec).MethodByName("Func1").Call(nil)
	fmt.Println("reflect.ValueOf(&myExec)", reflect.ValueOf(&myExec))
	fmt.Println("reflect.TypeOf(&myExec)", reflect.TypeOf(&myExec))
}

func(this *Execute)Func1(){
	fmt.Println("func1")
}
func (this *Execute)Func2(){
	fmt.Println("func2")
}
func(this *Execute)Func3(){
	fmt.Println("func3")
}