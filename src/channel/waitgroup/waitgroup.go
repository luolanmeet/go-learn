package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done func()
}

func doWork(id int, w worker) {

	for n := range w.in {
		fmt.Printf("worker %d received %c \n", id, n)
		w.done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func chanDemo() {

	var workers [10]worker
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)

	for i, worker := range workers {
		worker.in <- 'a' + i
		// add放在里边，有时候会报错
		// wg.Add(1)
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
		// wg.Add(1)
	}

	wg.Wait()
}

func main() {
	chanDemo()
}
