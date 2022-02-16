package main

import (
	"context"
	"time"
)

type otherContext struct {
	context.Context
}

func main() {
	// 使用 context.Background() 构建一个 WithCancel 类型的上下文
	ctxa, cancel := context.WithCancel(context.Background())

	// work模拟运行并检测前端退出通知
	go work(ctxa, "work-1")

	// 使用WithDeadline包装前面的上下文对象ctxa
	tm := time.Now().Add(3 * time.Second)
	ctxb, _ := context.WithDeadline(ctxa, tm)

	go work(ctxb, "work-2")

	// 使用WithValue包装前面的上下文对象ctxb
	oc := otherContext{ctxb}
	ctxc := context.WithValue(oc, "key", "pass from main")

	go workWithValue(ctxc, "work-3")

	// 睡眠10s，让work2，work3超时退出
	time.Sleep(10 * time.Second)

	// 显示调用work1的cancel()通知其退出，如果此时 work2和work3还没退出，同样也会被通知退出
	cancel()

	// 等待work1打印退出信息
	time.Sleep(5 * time.Second)
	println("main stop")
}

func workWithValue(ctx context.Context, name string) {

	for {
		select {
		case <-ctx.Done():
			println(name, " get msg to cancel")
			return
		default:
			value := ctx.Value("key").(string)
			println(name, " is running value=", value)
			time.Sleep(1 * time.Second)
		}
	}

}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			println(name, " get msg to cancel")
			return
		default:
			println(name, " is running")
			time.Sleep(1 * time.Second)
		}
	}
}
