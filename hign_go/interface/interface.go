package main

import (
	"fmt"
	"reflect"
)

type Animal interface {
	GetName() string
	PrintName()
}

type People struct {
	Name string
	Age  int
}

func (p *People) GetName() string {
	return p.Name
}
func (p *People) PrintName() {
	fmt.Println("name=", p.Name)
}

func main() {
	p1 := People{
		Name: "abcad",
		Age:  12,
	}
	var animal Animal
	animal = &p1
	fmt.Println(animal.GetName())
	animal.PrintName()

	fmt.Println("p1:", reflect.TypeOf(p1), "animal:", reflect.TypeOf(animal))
}
