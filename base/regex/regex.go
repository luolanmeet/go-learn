package main

import (
	"fmt"
	"regexp"
)

const text = `
my emial is mainem@163.com
he emial is mainem@qq.com.cn
`

func testSubMatch() {
	// 加()，可以提取需要的元素
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}

func main() {
	// [a-zA-Z0-9]@[a-zA-Z0-9].[a-zA-Z0-9.]
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9.]+`)
	//match := re.FindString(text)
	match := re.FindAllString(text, -1)
	fmt.Println(match)
	fmt.Println()
	testSubMatch()
}
