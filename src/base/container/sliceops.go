package main

import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
}

func define() {

	// slice 的zero value 是nil
	var s []int
	for i := 0; i < 50; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	// 创建了一个数组，返回其切片。len和cap会是3，并没有初始容量是要2的平方的限制
	s1 := []int{2, 4, 6}
	printSlice(s1)

	// 指定slice 的 array 的len/cap，array 的内容是对应类型的 zero value
	s2 := make([]int, 16)
	printSlice(s2)
	fmt.Println(s2)

	// array 的 len 和 cap 分开指定
	s3 := make([]int, 16, 32)
	printSlice(s3)
	fmt.Println(s3)

	// copy
	copy(s2, s1)
	printSlice(s2)
	fmt.Println(s2)

	// delete index=2 element
	s2 = append(s2[:2], s2[3:]...)
	printSlice(s2)
	fmt.Println(s2)

	fmt.Println("Poppint form front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println("Poppint from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(front, tail)
	fmt.Println(s2)
}

func main() {
	define()
}
