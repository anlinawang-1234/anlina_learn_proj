package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func main() {
	//MyErrFun()
	//check(1)
	//check(-1)
	//check(0)
	//if sentinelError() == sentinelErr{
	//	fmt.Println("equal error")
	//}
	//ErrorType()

	// 野生go routine
	GoRoutine(func() {
		fmt.Println("hello, world", reflect.TypeOf(errors.New("error")))
		panic("我异常了")
	})

	time.Sleep(10 * time.Second)
}

// === ErrorType
func ErrorType() {
	err := test()
	switch err := err.(type) {
	case nil:
		fmt.Println("no err")
	case *MyError:
		fmt.Println("MyErr", err.File, err.Line, err.Cnt)
	}
}

type MyError struct {
	Line string
	File string
	Cnt  int
}

func (e MyError) Error() string {
	return e.Line + e.File + strconv.Itoa(e.Cnt)
}
func test() error {
	return &MyError{"Line666", "File777", 89}
}

// === sentinelErr
var sentinelErr = errors.New("Sentinel Error")

func sentinelError() error {
	return sentinelErr
}

// 一般返回 bool error
func positive(n int) bool {
	if n == 0 {
		// bad case
		panic("undefined")
	}
	return n > -1
}

func check(n int) {
	defer func() {
		// 很不建议
		if recover() != nil {
			fmt.Println(n, "is neither")
		}
	}()
	if positive(n) {
		fmt.Println(n, "is positive")
	} else {
		fmt.Println(n, "is negative")
	}
}

// ==
type errorString string

func (e errorString) Error() string {
	return string(e)
}
func New(err string) error {
	return errorString(err)
}

var (
	MyErr  = New("myErr")
	LibErr = errors.New("libErr")
)

func MyErrFun() {
	fmt.Println(MyErr, "---", (New("myErr")))
	if MyErr == New("myErr") {
		fmt.Println("myErr")
	}

	fmt.Println(LibErr, "---", errors.New("libErr"))
	if LibErr == errors.New("libErr") {
		fmt.Println("libErr")
	}
}

// == 处理野生go routine的方法
func GoRoutine(x func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("被我抓到了", err, err.(reflect.Type))
			}
			fmt.Println("我要开始执行了")
			x()
		}()
	}()
}

// opaque error

type temp interface {
	Temp() bool
}

func isTemp(err error) bool {
	var e *temp
	// 断言
	if errors.As(err, e) {
		return err != nil && err.(temp).Temp()
	}
	return false
}
