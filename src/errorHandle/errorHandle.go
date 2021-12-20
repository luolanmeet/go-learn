package main

import (
	"errors"
	"fmt"
	"os"
)

func createError() {
	err := errors.New("this is my error. HA")
	fmt.Println(err)
}

func tryHandleError() {

	// 看方法注释，如果出现错误，返回的error是一个 PathError
	file, err := os.OpenFile("errorHandle.go", os.O_EXCL|os.O_CREATE, 0666)

	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s %s \n",
				pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}

	defer file.Close()
}

func main() {
	tryHandleError()
	createError()
}
