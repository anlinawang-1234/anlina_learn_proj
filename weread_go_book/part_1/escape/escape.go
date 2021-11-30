package main

func main() {
	escape2()
}

// 逃逸分析
var x *int

//go:noinline 增加这个注释表明地下的函数不进行内联调用
func escape1() {
	a := 1
	x = &a
}

var o *int

func escape2() {
	l := new(int) //  new(int) escapes to heap
	*l = 42
	m := &l
	n := &m
	o = **n
}
