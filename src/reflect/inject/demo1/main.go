package main

import (
	"fmt"
	"github.com/codegangsta/inject"
)

type S1 interface{}
type S2 interface{}

func Format(name string, company S1, level S2, age int) {
	fmt.Printf("name=%s company=%s level=%s age=%d", name, company, level, age)
}

func main() {
	// 控制实例的创建
	inj := inject.New()

	// 实参注入
	inj.Map("cck")
	inj.MapTo("alibaba", (*S1)(nil))
	inj.MapTo("p5", (*S2)(nil))
	inj.Map(27)

	// 函数反转调用
	inj.Invoke(Format)
}
