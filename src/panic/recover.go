package main

import (
	// "errors"
	"fmt"
)

func tryRecover() {

	// defer 执行了一个方法，最后的()其实是执行函数的意思， xxx()
	defer func() {
		// recover
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(fmt.Sprintf("can't handle. %v", r))
		}
	}()

	// 显示的报错
	// panic(errors.New("this is a error"))
	// panic(111)

	// 运行时的报错
	a := 0
	b := 1 / a
	fmt.Println(b)
}

func main() {
	tryRecover()
}
