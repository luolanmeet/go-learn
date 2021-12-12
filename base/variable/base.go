package main

// 导入的包必须使用，否则报错
import (
    "fmt"
    "math"
    "math/cmplx"
)

// 包内部变量
var count = 0

var (
    author = "cck"
    day = "2021年12月12日"
)

/**
 * 声明变量
 * 函数作用域的变量声明了就一定需要使用，否则会报错
 */
func variableDefine() {
    // 完整声明
    var a int = 1
    // 类型推断，省略声明
    var b = "hello"
    // 省略var关键字，只能在函数内使用
    c := true

    // 声明多个同类型的值
    var a1, a2, a3 int = 1, 2, 3
    // 声明多个不同类型的值
    var b1, b2, b3 = 1, true, "hh"
    // 简化写法
    c1, c2, c3 := 2, false, "hh"

    fmt.Println(a, b, c, a1, a2, a3, b1, b2, b3, c1, c2, c3)
}

/**
 * 变量零值 
 */ 
func variableZeroValue() {
    // 字符串
    var a string
    //(u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
    // 加 u 变为无符号整数，go没有long，需要大整数用 int64
    var b int 
    var c bool
    // float32 float64
    var d float32
    // 字符类型 不是char
    var e rune 
    // 复数类型 complex64(实部float32 虚部float32) complex128
    var f complex64
    var g error
    var h byte
    fmt.Println(a, b, c, d, e, f, g, h)
}


/**
 * 利用复数类型验证欧拉公式
 */
func euler() {
    fmt.Println(
        cmplx.Exp(1i * math.Pi) + 1,
        cmplx.Pow(math.E, 1i * math.Pi) + 1)
}

/**
 * 强制类型转换，需要显示地写出类型转换
 */ 
func forceTypeConvert() {
    var a, b int = 3, 4
    var c int
    c = int(math.Sqrt(float64(a * a + b * b)))
    fmt.Println(c)
}

/**
 * 常量
 * const 
 */
func consts() {

    // 不指定类型时可作为各种类型使用，某些情况下可以不用强制类型转换
    const a1, b1 = 3, 4
    var c1 int
    c1 = int(math.Sqrt(a1 * a1 + b1 * b1))
    fmt.Println(c1)

    // 指定了类型
    const (
        d1 int = 1
        e1 string = "hello"
    )
    fmt.Println(d1, e1)

    // 枚举，go没有枚举关键字，用一组const表示枚举，iota表示自增值
    const (
        cpp = iota
        java
        groovy
        _
        golang
    )
    fmt.Println(cpp, java, groovy, golang)

    // iota 结合表达式
    const (
        b = 1 << (10 * iota)
        kb
        mb
        gb
        tb
        pb
    )
    fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {
    fmt.Println(author, day)
    variableDefine()
    variableZeroValue()
    euler()
    forceTypeConvert()
    consts()
}