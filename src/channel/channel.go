package main

import (
	"fmt"
	"time"
)

func worker(c chan int) {

	for {
		fmt.Println("worker--")
		// 外部为channel赋值了，这里就会运行，否则会阻塞
		// ok 表示channel是否关闭
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("worker== %c \n", n)
	}

	// 另外的写法
	// for n := range c {
	// 	fmt.Printf("worker== %c \n", n)
	// }
}

func chanDemo() {
	c := make(chan int)
	go worker(c)
	// 发送了数据，但没有人收，也会阻塞
	fmt.Println("chanDemo-- ")
	c <- 1
	// c <- 2
	fmt.Println("chanDemo==")
	time.Sleep(time.Millisecond)
}

/**
 * 返回一个channel，只要往这个channel传数据，这个worker就会工作
 * chan<- 表示返回的channel，只能用来传数据，不能用了收数据
 */
func createWorker(id int) chan<- int {

	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received ceived %c\n", id, <-c)
		}
	}()

	return c
}

func bufferChannel() {
	// 缓冲区大小为3，缓冲区未满之前，写数据，都不会阻塞
	c := make(chan int, 3)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	go worker(c)

	// 关闭channel，关闭后，接受方会一直接收到空数据
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	// chanDemo()
	// c := createWorker(1)
	// c <- 'A'
	// time.Sleep(time.Millisecond)
	bufferChannel()

}
