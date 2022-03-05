package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	u := User{1, "cck", 28}
	va := reflect.ValueOf(u)  // 值类型
	vb := reflect.ValueOf(&u) // 指针类型

	fmt.Println(va.CanSet(), va.FieldByName("Name").CanSet())
	fmt.Println(vb.CanSet(), vb.Elem().FieldByName("Name").CanSet())

	fmt.Printf("%v\n", vb)
	name := "ck"
	vc := reflect.ValueOf(name)
	// 通过Set函数修改变量的值
	vb.Elem().FieldByName("Name").Set(vc) // 没法直接 Set(name)，Set的入参是一个Value
	fmt.Printf("%v\n", vb)
}
