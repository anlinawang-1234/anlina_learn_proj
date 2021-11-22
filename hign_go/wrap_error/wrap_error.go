package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

func main() {
	config, err := ReadConfig("./config.txt")
	if err != nil {
		fmt.Printf("cause %T --- %+v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace\n %+v\n", err)
		os.Exit(1)
	}
	fmt.Println("读取成功", string(config))
}

func ReadConfig(path string) ([]byte, error) {
	config, err := ReadAll(path)
	if err != nil {
		return []byte{}, errors.WithMessage(err, "read config error")
	}
	return config, nil
}

func ReadAll(path string) ([]byte, error) {
	tmpFile, err := os.Open(path)
	if err != nil {
		return []byte{}, errors.Wrap(err, "open file err")
	}

	tmpByte, err := ioutil.ReadAll(tmpFile)
	if err != nil {
		return []byte{}, errors.Wrap(err, "read file err")
	}
	return tmpByte, nil
}
