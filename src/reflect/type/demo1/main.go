package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string "学生姓名"
	Age  int    `a:"111"b:"222"`
}

func main() {
	s := Student{Name: "cck"}
	rt := reflect.TypeOf(s)

	fieldName, ok := rt.FieldByName("Name")
	// 取Tag数据
	if ok {
		fmt.Println(fieldName.Tag)
	}
	fieldAge, ok := rt.FieldByName("Age")
	if ok {
		fmt.Println(fieldAge.Tag.Get("a"))
		fmt.Println(fieldAge.Tag.Get("b"))
	}

	fmt.Println("typeName:", rt.Name())
	fmt.Println("NumField:", rt.NumField())
	fmt.Println("PkgPath:", rt.PkgPath())
	fmt.Println("String:", rt.String())
	fmt.Println("Kind.String:", rt.Kind().String())

	// 获取结构类型的字段名称
	for i := 0; i < rt.NumField(); i++ {
		fmt.Printf("type.Field[%d].Name = %v \n", i, rt.Field(i).Name)
	}

	sc := make([]int, 10)
	sc = append(sc, 1, 2, 3)
	sct := reflect.TypeOf(sc)
	// 获取slice元素的Type
	scet := sct.Elem()

	fmt.Println("slice element type.Kind()=", scet.Kind())
	fmt.Printf("slice element type.Kind()=%d \n", scet.Kind())
	fmt.Println("slice element type.String()=", scet.String())

	fmt.Println("slice element type.Name()=", scet.Name())
	fmt.Println("slice element type.NumMethod()=", scet.NumMethod())
	fmt.Println("slice element type.PkgPath()=", scet.PkgPath())
	fmt.Println("slice element type.PkgPath()=", sct.PkgPath())

}
