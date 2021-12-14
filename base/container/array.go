package main

import (
    "fmt"
)

func define() {

    var arr1 [5]int
    // 用 := 必须给定值
    arr2 := [3]int{1, 2, 3}
    // 用...可以省略数组长度
    arr3 := [...]int{1, 2}
    // 3行4列
    var arr4 [3][4]bool

    fmt.Println(arr1, arr2, arr3)
    fmt.Println(arr4)
}

func traversal() {

    arr := [...]int{1, 2}

    // 普通遍历
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }
    // 利用range获取下标遍历
    for i := range arr {
        fmt.Println(arr[i])
    }
    // 遍历下标和值
    for i, v := range arr {
        fmt.Println(i, v)
    }
    // 遍历数组的值
    for _, v := range arr {
        fmt.Println(v)
    }

}

func printArray(arr *[5]int) {

    // 这里不用写成 (*arr)[0]
    fmt.Println(arr[0])

    for i, v := range arr {
        fmt.Println(i, v)
    }
}

func main() {

    define()
    traversal()

    // 数组也是值类型，调用函数是传入数组会发生数组拷贝
    // [5]int 和 [3]int 是不一样的类型！
    // go 语言中一般不直接使用数组，而是切片
    arr := [5]int{1, 2, 3, 4, 5}
    printArray(&arr)
    // arr1 := [3]int{1, 2, 3}
    // printArray(&arr1)

}