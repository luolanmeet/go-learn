package main

import (
	"os"
	"text/template"
)

func t1() {
	// 数据
	name := "cck"
	// 定义模版
	muban := "hello, {{.}}"
	// 解析模版
	tmpl, err := template.New("t1").Parse(muban)
	if err != nil {
		return
	}
	// 执行目标，输出到终端
	err = tmpl.Execute(os.Stdout, name)
	if err != nil {
		return
	}
}

type Person struct {
	Name string
	Age  int
}

// 使用结构体
func t2() {

	// 数据
	cck := Person{
		Name: "cck",
		Age:  50,
	}

	// 定义模版
	muban := "hello, {{.Name}}, your age {{.Age}}"
	// 解析模版
	tmpl, _ := template.New("t2").Parse(muban)
	// 执行目标，输出到终端
	tmpl.Execute(os.Stdout, cck)
}

func main() {
	//t1()
	t2()
}
