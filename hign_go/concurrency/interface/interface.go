package main

import (
	"fmt"
	"reflect"
)

type IceCream interface {
	Speak()
}

type Ben struct {
	Name string
}

func (b *Ben) Speak() {
	fmt.Println("I am ", b.Name)
}

type Jerry struct {
	Id   int32
	Name string
}

func (j *Jerry) Speak() {
	fmt.Println(j.Name, "我是", j.Name, j.Id)
}

func main() {
	ben := Ben{Name: "bennnn"}
	jerry := Jerry{Name: "jerry", Id: 666}
	ice := IceCream(&ben)
	var loop0, loop1 func()
	loop0 = func() {
		ice = &ben
		go loop1()
	}

	loop1 = func() {
		ice = &jerry
		go loop0()
	}
	go loop0()

	for {
		ice.Speak()
		fmt.Println("ice", reflect.TypeOf(ice), reflect.ValueOf(ice))
	}
	//ice.Speak()
	//
}
