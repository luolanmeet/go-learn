package main

import (
    "fmt"
)

// [] 表示是一个切面，而不是一个数组
func updateArr(s []int) {
    s[0] = 100
}
 
func base() {
    arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
    // 半开半闭
    fmt.Println(arr[2:6])
    fmt.Println(arr[:6])
    fmt.Println(arr[2:])
    fmt.Println(arr[:])

    // 切面slice 是数组的 一个视图
    // slice 本身没有数据

    // 更新了 arr[0]
    updateArr(arr[:])
    fmt.Println(arr)

    // 更新了 arr[2]
    updateArr(arr[2:])
    fmt.Println(arr)

    // reslice
    s := arr[:]
    s = s[:5]
    s = s[2:]
    fmt.Println(s)
}

func extend() {

    // slice -> (ptr, len, cap)
    // slice 可以往后扩展，不可以向前扩展
    // s[i] 不可以跨越len，向后扩展不可以超过cap 

    arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
    s1 := arr[2:6]
    s2 := s1[3:5]
    fmt.Println(s1)
    fmt.Println(s2)

    fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d \n", s1, len(s1), cap(s1))
}

func main() {
    base()
    extend()
}
