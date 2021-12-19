package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0

	// 对于返回出去的 func，sum变量是他的上下文变量的
	return func(v int) int {
		sum += v
		return sum
	}
}

func addFunc(val int) int {
	sum := 0
	sum += val
	return sum
}

func main() {

	// 发现 sum 这变量，是可以完成累加的
	// 闭包的概念
	a := adder()
	for i := 1; i < 10; i++ {
		fmt.Printf("0 + ... + %d = %d\n", i, a(i))
	}

	// 这样并不能完成累计，sum每次都重置为是0，即便是用了同个函数变量
	add := addFunc
	fmt.Printf("%T %v\n", add, add)
	fmt.Println(add(1))
	fmt.Println(add(1))
	fmt.Println(add(1))

}
