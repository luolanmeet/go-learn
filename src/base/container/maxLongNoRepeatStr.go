package main

import (
	"fmt"
)

func maxLongNoRepeatStr(str string) string {

	m := make(map[rune]int)

	start, maxLen, left, right := 0, 0, 0, 0

	runs := []rune(str)

	for i, ch := range runs {

		if lastIdx, ok := m[ch]; ok && lastIdx >= start {
			start = lastIdx + 1
		}

		if i-start+1 > maxLen {
			maxLen = i - start + 1
			left = start
			right = start + maxLen
		}
		m[ch] = i
	}

	return string(runs[left:right])
}

func main() {

	fmt.Println(maxLongNoRepeatStr("helloworld!"))
	fmt.Println(maxLongNoRepeatStr("你好哇世界哈哈哈!"))

}
