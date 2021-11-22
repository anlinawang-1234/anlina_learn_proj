package main

func main() {

}

type MyError interface {
	Error() string
}

type TmpError struct {
}
