package main

import (
	"fmt"
)

// [] 表示是一个切面，而不是一个数组
func updateArr(s []int) {
	s[0] = 100
}

func base() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	// 半开半闭
	fmt.Println(arr[2:6])
	fmt.Println(arr[:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:])

	// 切面slice 是数组的 一个视图
	// slice 本身没有数据

	// 更新了 arr[0]
	updateArr(arr[:])
	fmt.Println(arr)

	// 更新了 arr[2]
	updateArr(arr[2:])
	fmt.Println(arr)

	// reslice
	s := arr[:]
	s = s[:5]
	s = s[2:]
	fmt.Println(s)
}

func extend() {

	// slice -> (ptr, len, cap)
	// slice 可以往后扩展，不可以向前扩展
	// s[i] 不可以跨越len，向后扩展不可以超过cap

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d \n", s1, len(s1), cap(s1))
}

func testAppend() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	// 想slice添加元素时如果超越cap，系统会重新分配更大的底层数组
	// 由于值传递的关系，必须接受append的返回值
	fmt.Println(s3, s4, s5)
	fmt.Println("arr = ", arr)
}

func other() {
	// s1 切面内部指向的是不是原数组。是原数组
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[0:2]
	fmt.Printf("%v \n", s1)
	arr[0] = 100
	s1[1] = 200
	fmt.Printf("%v \n", s1)

	// make 创建的切片
	s2 := make([]int, 2)
	// 分隔切片后，新的切片仍和旧的切片指向同个数组
	s3 := s2[:]
	fmt.Printf("%v \n", s3)
	s2[0] = 100
	s2[1] = 100
	fmt.Printf("%v \n", s3)
}

func testAppend2() []int {
	s1 := make([]int, 3)
	s1[0] = 0
	s1[1] = 1
	s1[2] = 2

	// append 之后生成的slice，指向的是同个数组吗
	// 不是！
	s2 := append(s1, 3)
	fmt.Printf("%v \n", s2)
	s1[0] = 100
	fmt.Printf("%v \n", s2)

	// 由切面产生的切面，指向的是同个数组吗
	// 是
	s3 := s1[0:1]
	fmt.Printf("%v \n", s3)
	s1[0] = 1000
	fmt.Printf("%v \n", s3)

	return s1
}

func testAppend3() {

	s := make([]int, 3, 3)
	s1 := append(s, 1)
	fmt.Printf("%v \n", s)
	fmt.Printf("%v \n", s1)
}

func main() {
	fmt.Println("base", "=============")
	base()

	fmt.Println("extend", "=============")
	extend()

	fmt.Println("testAppend", "=============")
	testAppend()

	fmt.Println("other", "=============")
	other()

	fmt.Println("testAppend2", "=============")
	// 切面不能作为返回值？可以啊
	s := testAppend2()
	fmt.Println(s)

	fmt.Println("testAppend3", "=============")
	testAppend3()

}
