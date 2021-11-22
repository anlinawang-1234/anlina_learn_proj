package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *People) Print() {
	fmt.Println("[people]", fmt.Sprintf("%s_%d", p.Name, p.Age))
}

type Student struct {
	People
	Name string
	Sex  int
}

//func(s *Student)Print(){
//	fmt.Println("[Student]", fmt.Sprintf("%s_%d_%s_%d", s.People.Name, s.People.Age, s.Name,s.Sex))
//}
func main() {
	// 获取tag
	TypeOfPeople := reflect.TypeOf(People{})
	for i := 0; i < TypeOfPeople.NumField(); i++ {
		fmt.Println("filed", TypeOfPeople.Field(i), "tag", TypeOfPeople.Field(i).Tag.Get("json"))
	}

	// 字段重复先用子类
	stu := Student{
		People: People{
			Name: "abc",
			Age:  15,
		},
		Name: "efadfd",
		Sex:  1,
	}
	fmt.Println("std", stu, stu.Name)
	stu.Print()
}
