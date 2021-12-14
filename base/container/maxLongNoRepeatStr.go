package main

import(
    "fmt"
)

func maxLongNoRepeatStr(str string) string {

    m := make(map[byte]int)

    start, maxLen, left, right := 0, 0, 0, 0

    for i, ch := range []byte(str) {

        if lastIdx, ok := m[ch]; ok && lastIdx >= start {
            start = lastIdx + 1
        }
        
        if i - start + 1 > maxLen {
            maxLen = i - start + 1
            left = start
            right = start + maxLen
        }
        m[ch] = i
    }

    return str[left:right]
}


func main() {

    fmt.Println(maxLongNoRepeatStr("helloworld!"))

}