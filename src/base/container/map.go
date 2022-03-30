package main

import (
	"fmt"
)

func baseFunc() {

	// 最后一个值，需要加","号，或者是和}放同一行
	m1 := map[string]string{
		"name":   "cck",
		"coutry": "china",
	}

	//  empty map
	m2 := make(map[string]int)

	// nil
	var m3 map[string]int

	fmt.Println(m1, m2, m3)
	fmt.Println(len(m1))

	// 遍历 是无序的
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// map使用哈希表，必须可以比较相等
	// 除了 slice,map,function的内建类型，其他都可以作为map的key
	// Struct类型不包含上述字段时，也可以作为key，编译时会自动检查

}

func ops() {

	m := map[string]string{
		"name":   "cck",
		"coutry": "china",
	}

	// m[key] 获取元素
	name := m["name"]
	fmt.Println(name)
	// key不存在时，获取value类型的零值
	name1 := m["name1"]
	fmt.Println(name1)

	// 可以用第二个值来判断key是否存在
	coutry, ok := m["coutry"]
	fmt.Println(coutry, ok)

	if coutry1, ok := m["coutry1"]; ok {
		fmt.Println(coutry1)
	} else {
		fmt.Println("key does not exist")
	}

	// 删除元素
	delete(m, "name")
}

func main() {

	// map[k]v

	baseFunc()
	ops()

}
