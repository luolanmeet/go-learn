package main

import (
	"math/rand"
)

func GenerateInt(done chan struct{}) chan int {

	ch := make(chan int, 10)

	go func() {
	LABEL:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break LABEL
			}
		}
		close(ch)
	}()

	return ch
}

func GenerateInt2(done chan struct{}) chan int {

	ch := make(chan int, 20)
	send := make(chan struct{})
	go func() {

	LABEl:
		for {
			select {
			case ch <- <-GenerateInt(send):
			case ch <- <-GenerateInt(send):
			case <-done:
				break LABEl
			}
		}
		close(ch)
	}()

	return ch
}

func main() {

	done := make(chan struct{})

	//ch := GenerateInt(done)
	ch := GenerateInt2(done)

	for i := 0; i < 10; i++ {
		println(<-ch)
	}

	// 通知停止生成随机数
	done <- struct{}{}

	println(<-ch)

}
