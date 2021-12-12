package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "os"
    "bufio"
)

func iffunc() {

    // 条件不需要用()包裹
    var a = 1
    if a > 0 {
    } else {
    }

    // 可以先执行语句，再进行判断
    const filename = "abc.txt"
    if contents, err := ioutil.ReadFile(filename); err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%s\n", contents)
    }

    // 出了if，就不能访问到if的变量了 
    // fmt.Printf("%s\n", contents)

}

func switchfunc() {

    // case 后不需要有break，默认就有
    // 需要不break时，加fallthrough

    // 常规使用 switch xx {case xxx}
    a := 3
    switch a {
        case 1:
            fmt.Println("一")
        case 2:
            fmt.Println("二")
        case 3:
            fmt.Println("三")
            fallthrough
        default:
            fmt.Println("未知")
    }

    // switch {case 表达式} 
    score := 80
    switch {
    case score < 0 || score > 100:
        panic(fmt.Sprintf("Wrong score: %d", score))
    case score < 60:
        fmt.Println("D")
    case score < 80:
        fmt.Println("C")
    case score < 90:
        fmt.Println("B")
    case score <= 100:
        fmt.Println("1") 
    }

}

func forfunc1(num int) string {
    // 正常的for循环
    res := ""
    for ; num != 0; num /= 2 {
        res = strconv.Itoa(num % 2) + res
    }
    return res
}

func forfunc2() {

    filename := "abc.txt"
    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)

    // 类似while的for循环
    for scanner.Scan() {
        fmt.Println(scanner.Text())        
    }

    // 死循环
    // for { 
    // }
}

func main() {

    iffunc()
    switchfunc()
    fmt.Println(forfunc1(13))
    forfunc2()

}