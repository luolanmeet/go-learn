package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var a []int         // nil 切片
	b := make([]int, 0) // 空切片

	if a == nil {
		fmt.Println("a is nil")
	} else {
		fmt.Println("a is not nil")
	}

	// b的底层数组大小为0，但切片不是nil
	if b == nil {
		fmt.Println("b is nil")
	} else {
		fmt.Println("b is not nil")
	}

	// 使用反射中的SliceHeader来获取切片运行时的数据结构
	as := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bs := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	fmt.Printf("len=%d,cap=%d,type=%d\n", len(a), cap(a), as.Data)
	fmt.Printf("len=%d,cap=%d,type=%d\n", len(b), cap(b), bs.Data)

}
