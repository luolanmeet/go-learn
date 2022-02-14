package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {

	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()

	return out
}

func main() {

	var c1, c2 = generator(), generator()

	// 这是一个nil channel
	// nil channel在select 中是不会报错的，但也不会被执行
	var c3 chan int

	// 返回的是一个channel ， 10秒钟后，会往tm这个chanel中传值
	tm := time.After(10 * time.Second)

	// 每秒钟产生一个数据
	tick := time.Tick(time.Second)

	for {
		select {
		case n := <-c1:
			fmt.Println("Received from c1: ", n)
		case n := <-c2:
			fmt.Println("Received from c2: ", n)

		// c3 是个nilchannel
		case n := <-c3:
			fmt.Println("Received from c3: ", n)

		// 每次select 超过指定时间没有数据，则打印timeout
		case <-time.After(600 * time.Millisecond):
			fmt.Println("timeout")

		case <-tick:
			fmt.Println("tick")

		case n := <-tm:
			fmt.Println("Bye ", n)
			return
			// default:
			// 	fmt.Println("No value received")

		}
	}

}
