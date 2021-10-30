package main


type S struct {}
func main() {
	var x S
	_ = identity(x)
}
func identity(x S) S {
	return x
}

// 逃逸分析
//func foo()*int{
//	t := 3
//	return &t
//}
//
//func main(){
//	//x := foo()
//	//fmt.Println("x", *x)
//}
