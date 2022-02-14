package main

func main() {

	//ch := make(chan int, 1)
	ch := make(chan int)
	go func(chan int) {

		for {
			// 有多个可用时，select随机选取一个处理
			select {
			case ch <- 10:
			case ch <- 20:
			}
		}

	}(ch)

	for i := 0; i < 10; i++ {
		println(<-ch)
	}

	//for v := range ch {
	//	println(v)
	//}

}
