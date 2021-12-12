package main

import (
    "fmt"
)

func base() {

    // 指针不能运算
    var a int = 2
    var pa *int = &a
    *pa = 3
    fmt.Println(a)
}

func swap(a, b *int) {
    *a, *b = *b, *a
}

func main() {

    // 值传递，传递参数时，复制一份数据
    // 引用传递，传递参数时，传入数据的引用，实际是同一份数据

    // Go只有值传递一种方式，但因为有指针，所以也可以避免大数据做值传递时的性能问题，实现引用传递
    base()

    a, b := 3, 4
    swap(&a, &b)
    fmt.Println(a, b)
}