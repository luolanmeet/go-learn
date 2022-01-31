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

	// 值共享
	// 期望是 b = [1,3] c = [2,4]
	// 结果是 b = [1,3] c = [3,4]
	a := []int{1, 2}
	fmt.Printf("a =  %v \n", a)

	// 这里，append 之后，把slice a的值也给改了
	b := append(a[0:1], 3)
	fmt.Printf("a =  %v \n", a)

	c := append(a[1:2], 4)
	fmt.Printf("b =  %v \n", b)
	fmt.Printf("c =  %v \n", c)

	// 可以强迫追加到新数组
	fmt.Println("--------")
	a = []int{1, 2}
	fmt.Printf("a =  %v \n", a)

	b = append(a[0:1:1], 3)
	fmt.Printf("a =  %v \n", a)

	c = append(a[1:2:2], 4)
	fmt.Printf("b =  %v \n", b)
	fmt.Printf("c =  %v \n", c)

}

func testAppend4() {

	// len 是 0 ，所以初始值 {0}，cap是2
	s := make([]int, 1, 2)
	fmt.Printf("s=%v \n", s)

	// append 的时候cap够用，不会新开内存空间，而是用原来的s指向的数组存储 {0,1}
	// append 之后，s的cap还是不会变
	s1 := append(s, 1)
	fmt.Printf("s=%v len=%d cap=%d \n", s, len(s), cap(s))
	fmt.Printf("s1=%v \n", s1)

	// append 的时候cap够用，不会新开内存空间，而是用原来的s指向的数组存储 {0,2}
	// 此时由于s1和s2都是指向同个数组，因此s1的值也会被改变{0,1} -> {0,2}
	s2 := append(s, 2)
	fmt.Printf("s1=%v \n", s1)
	fmt.Printf("s2=%v \n", s2)
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

	fmt.Println("testAppend4", "=============")
	testAppend4()

}
