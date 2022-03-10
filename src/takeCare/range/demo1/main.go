package main

import (
	"sync"
	"time"
)

func test1() {
	wg := sync.WaitGroup{}
	nums := []int{1, 2, 3, 4, 5}
	for i := range nums {
		wg.Add(1)
		go func() {
			// i会被复用，i的值是固定的吗？不是的
			// main的goroutine和新开的goroutine对于i存在数据竞争。新开的goroutine获取到i是不确定的
			println(nums[i])
			wg.Done()
		}()

		time.Sleep(1 * time.Millisecond)
		//time.Sleep(2 * time.Second)
	}

	wg.Wait()
}

func test2() {
	wg := sync.WaitGroup{}
	nums := []int{1, 2, 3, 4, 5}
	for i := range nums {
		wg.Add(1)
		// 实参到形参的值拷贝
		go func(idx int) {
			println(nums[idx])
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {

	//  go run -race src/takeCare/range/demo1/main.go
	test1()
	println()
	test2()
}
