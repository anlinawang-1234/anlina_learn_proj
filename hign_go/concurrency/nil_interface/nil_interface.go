package main

import (
	"fmt"
	"reflect"
)

type IPeople interface {
	hello()
}
type Jerry struct{}

func (j *Jerry) hello() {
	fmt.Println("I am jerry")
}
func errFunc(i int) *Jerry {
	if i == 0 {
		fmt.Println("内部返回nil")
		return nil
	} else {
		fmt.Println("内部返回数组")
		return &Jerry{}
	}
}

func main() {
	var i IPeople
	// 第一版
	i = errFunc(0)
	if i == nil {
		fmt.Println("外部返回的也是nil")
	} else {
		// 输出在这里
		fmt.Println("出错了，外部返回的不是nil")
		// type=*main.Jerry value=<nil> interface两个指针，一个指向值一个执行类型，只有两个都为nil时，interface才为nil
		fmt.Printf("i %v %+v type=%+v value=%+v\n", i, &i,
			reflect.TypeOf(i), reflect.ValueOf(i))
	}

	// 改进
	//out := errFunc(1)
	//if out != nil{
	//	i = out
	//	fmt.Println("这里ok了", i)
	//}else{
	//	fmt.Println("外部返回的也是nil")
	//}
}
