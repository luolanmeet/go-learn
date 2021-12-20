package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	// 先声明的defer后执行
	// 栈 先进后出
	defer fmt.Println(1)
	defer fmt.Println(2)
	// panic("error occurred")
	return
	// 报错或return 都无法被执行
	defer fmt.Println(3)
	fmt.Println(4)

}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}

}

func writeFile(fileName string) {

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	// 申请了资源，就应该立即想到释放资源
	defer file.Close()

	// 用了buffer，就应该立即想到刷新buffer到磁盘
	write := bufio.NewWriter(file)
	defer write.Flush()

	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(write, f())
	}
}

func main() {
	tryDefer()
	writeFile("fibonacci.txt")
}
