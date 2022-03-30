package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "gogo程序员"
	fmt.Println(s)

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	// 使用库
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// ch 是一个 rune
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	// 直接转rune，是重新编码，开辟了个rune数组的空间存储转换结果，每个rune4个字节
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()

}
