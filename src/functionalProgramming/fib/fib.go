package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}

}

type intGen func() int

func fibonacci2() intGen {
	a, b := 0, 1
	// intGen 就是一个func，所以这样返回没有问题
	return func() int {
		a, b = b, a+b
		return a
	}

}

/**
 * 函数也能实现接口，真是秀
 */
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()

	if next > 10000 {
		return 0, io.EOF
	}

	s := fmt.Sprintf("%d\n", next)
	// TODO p 太小会有问题
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {

	fmt.Println("fibonacci")
	f := fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("fibonacci2")
	f2 := fibonacci2()
	printFileContents(f2)
}
