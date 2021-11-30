package main

// 生成ssa文件的命令: GOSSAFUNC=main GOSS=linux GOARCH=amd64 go tool compile ssa.go
var d int

func main() {
	a := 1
	if true {
		a = 2
	}
	d = a
}
