package main

import "time"

// 一个查询结构体
type query struct {
	// 参数channel
	sql chan string
	// 结果channel
	result chan string
}

func execQuery(q query) {

	// 启动协程
	go func() {
		// 获取输入
		sql := <-q.sql
		// 访问数据库
		// 输出结果
		q.result <- "result form " + sql
	}()
}

func main() {

	// 初始化查询
	q := query{sql: make(chan string, 1), result: make(chan string, 1)}
	// 这里需要用 go execQuery(q) 吗？
	execQuery(q)

	// 发送参数
	q.sql <- "select * from table"

	// 其他业务逻辑
	println("其他业务逻辑...")
	time.Sleep(1 * time.Second)

	// 获取结果
	println(<-q.result)
}
