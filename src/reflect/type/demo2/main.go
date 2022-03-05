package main

import "reflect"

type INT int

type A struct {
	a int
}

type B struct {
	b string
}

type Ita interface {
	String() string
}

func (b B) String() string {
	return b.b
}

func main() {
	var a INT = 12
	var b int = 14

	// 实参是具体类型，reflect.TypeOf返回的是其静态类型
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	// INT 和 int 是两个类型，两者并不相同
	if ta == tb {
		println("ta==tb")
	} else {
		println("ta!=tb")
	}
	println(ta.Name()) // INT
	println(tb.Name()) // int
	// 底层是基础类型
	println(ta.Kind().String()) // int
	println(tb.Kind().String()) // int

	s1 := A{1}
	s2 := B{"tt"}
	// 实参是具体类型，reflect.TypeOf返回的是其静态类型
	ts1 := reflect.TypeOf(s1)
	ts2 := reflect.TypeOf(s2)
	println(ts1.Name()) // A
	println(ts2.Name()) // B

	// Kind返回的是基本类型
	println(ts1.Kind().String()) // struct
	println(ts2.Kind().String()) // struct

	ita := new(Ita)
	var itb Ita = s2

	// 实参是未绑定具体变量的接口类型，reflect.TypeOf 返回的是接口类型本身
	println(reflect.TypeOf(ita).Elem().Name())          // Ita
	println(reflect.TypeOf(ita).Elem().Kind().String()) // interface

	// 实参是绑定了具体变量的接口类型，reflect.TypeOf 返回的是绑定的具体类型
	// 如何知道是否进行了绑定？
	println(reflect.TypeOf(itb).Name())          // B
	println(reflect.TypeOf(itb).Kind().String()) // struct

}
