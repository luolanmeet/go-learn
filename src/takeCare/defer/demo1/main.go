package main

import "fmt"

// 带命名返回值函数
func f1() (r int) {

	// 函数调用方负责开辟空间，包括形参和返回值的空间
	// 有名的函数返回值相当于函数的局部变量，r分配在栈上，其地址被称为 返回值栈区，初始化为零值 -> 0，
	// return 10 ，10会复制到返回值栈区，r变为10
	// defer 这里相当于对函数返回值r的闭包引用，r++之后，变为11
	defer func() {
		r++
	}()
	return 10
}

func f2() (r int) {

	// r分配在栈上，其地址被称为 返回值栈区，初始化为零值 -> 0，
	// 引入局部变量t，赋值为5
	// return t,赋值t的值5到返回值r所在的栈区，r变为5（注意！这里是值拷贝，并不是t就是返回值变量了，r才是）
	// defer 是对变量t的闭包引用，修改t之后t边为10，但返回值栈区的值还是5
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {

	// r分配在栈上，其地址被称为 返回值栈区，初始化为零值 -> 0，
	// return 1，复制1到函数返回值r所在的栈区

	// defer这里，是传值，设计值复制，已经不是同个变量了
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func f4() int {

	// 声明并初始化了局部变量 r
	// return r， r被 值拷贝 到 返回值栈区，返回值栈区值为1
	// defer对局部变量++，r变为1

	// 这里没使用有名返回值函数，从根本上避免了 defer 对返回值栈区的引用

	r := 1
	defer func() {
		r++
	}()
	return r
}

func main() {
	fmt.Println("f1=", f1()) // 11
	fmt.Println("f2=", f2()) // 5
	fmt.Println("f3=", f3()) // 1
	fmt.Println("f4=", f4()) // 1
}
