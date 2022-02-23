package main

import (
	"sync"
	"sync/atomic"
)

var total uint32

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	var i uint32
	for i = 0; i < 100; i++ {
		atomic.AddUint32(&total, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()

	println("total value=", total)
}
