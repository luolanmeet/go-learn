package main

import "runtime"

func main() {

	c := make(chan struct{})
	ci := make(chan int, 10)

	go func(i chan struct{}, j chan int) {

		for i := 1; i <= 10; i++ {
			ci <- i
		}
		// 这里关闭ci，下边for循环才不会有死锁
		close(ci)
		c <- struct{}{}
	}(c, ci)

	println("NumGoroutine", runtime.NumGoroutine())
	<-c
	println("NumGoroutine", runtime.NumGoroutine())

	// ci 已经关闭，但还是可以读取到数据
	for v := range ci {
		println(v)
	}

	// 读到零值
	println("====")
	println(<-ci)

}
