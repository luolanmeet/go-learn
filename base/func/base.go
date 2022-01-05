package main

import (
	"fmt"
)

func div0(a, b int) int {
	return a / b
}

/**
 * 可以返回多个值
 */
func div1(a, b int) (int, int) {
	return a / b, a % b
}

/**
 * 可以指定返回值的名字
 */
func div2(a, b int) (q, r int) {
	// 下边的计算注释掉也可以正常运行，返回的是qr的零值
	q = a / b
	r = a % b
	return
}

/**
 * 参数传入函数
 * 函数式编程
 */
func apply(op func(int, int) int, a, b int) int {
	return op(a, b)
}

/**
 * 可变参数列表
 */
func sum(numbers ...int) int {

	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func main() {

	fmt.Println(div1(5, 2))
	fmt.Println(div2(5, 2))

	fmt.Println(apply(div0, 5, 2))
	fmt.Println(apply(func(a, b int) int {
		return a * b
	}, 5, 2))

	fmt.Println(sum(1, 2, 3))

}
