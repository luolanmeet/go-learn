package main

import (
	"fmt"
)

type worker struct {
	in   chan int
	done chan bool
}

func doWork(id int, c chan int, done chan bool) {
	for {
		fmt.Printf("worker %d received %c \n", id, <-c)

		go func() {
			done <- true
		}()
	}
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}

func chanDemo() {
	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
		// 当能收到done的数据时，代表in的数据已经被处理好了
		// 如果打开这段代码，就变成串行执行了
		// <-worker.done
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// wait for all done
	// 用channel 来通信
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}

}

func main() {
	chanDemo()
}
