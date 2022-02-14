package main

import "sync"

// 定时任务结构体
type task struct {
	begin int
	end   int
	// 只发送不接收的chan
	result chan<- int
}

// 计算开始到结束的自然数之和，结果放入到result通道中
func (t *task) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum
}

func main() {

	// 创建任务通道
	taskchan := make(chan task, 10)

	// 创建结果通道
	resultchan := make(chan int, 10)

	// wait用于同步等待任务的执行
	wait := &sync.WaitGroup{}

	// 初始化task的goroutine，计算100个自然数之和
	go InitTask(taskchan, resultchan, 100)

	// 每个task启动一个goroutine进行处理
	go DistributeTask(taskchan, wait, resultchan)

	// 通过结果通道获取结果并汇总
	sum := ProcessResult(resultchan)

	println("sum=", sum)
}

func ProcessResult(resultchan chan int) int {

	sum := 0
	for v := range resultchan {
		sum += v
	}

	return sum
}

// 读取task chan，每个task启动一个worker goroutin进行处理
// 并等待每个task运行完，关闭结果通道
func DistributeTask(taskchan <-chan task, wait *sync.WaitGroup, resultchan chan int) {

	for v := range taskchan {
		wait.Add(1)
		go ProcessTask(v, wait)
	}

	wait.Wait()
	close(resultchan)
}

// goroutine 处理具体工作，并将处理结果发送到结果通道
func ProcessTask(t task, wait *sync.WaitGroup) {
	t.do()
	wait.Done()
}

// 构建task并写入task通道
func InitTask(taskchan chan<- task, resultchan chan int, p int) {

	qu := p / 10
	mod := p % 10
	high := qu * 10
	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)
		task := task{
			begin:  b,
			end:    e,
			result: resultchan}
		taskchan <- task
	}
	if mod != 0 {
		task := task{
			begin:  high + 1,
			end:    p,
			result: resultchan}
		taskchan <- task
	}
	close(taskchan)
}
