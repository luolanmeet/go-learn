package main

// 工作池的goroutine数目
const (
	NUMBER = 10
)

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

	workers := NUMBER

	// 创建任务通道
	taskchan := make(chan task, 10)

	// 创建结果通道
	resultchan := make(chan int, 10)

	// worker信号通道
	done := make(chan struct{}, 10)

	// 初始化task的goroutine，计算100个自然数之和
	go InitTask(taskchan, resultchan, 100)

	// 分配任务到NUMBER个goroutine
	DistributeTask(taskchan, workers, done)

	// 获取各个goroutine处理完任务的通知，并关闭结果通道
	go CloseResult(done, resultchan, workers)

	// 通过结果通道获取结果并汇总
	sum := ProcessResult(resultchan)

	println("sum=", sum)
}

// 读取结果通道，汇总数据
func ProcessResult(resultchan chan int) int {
	sum := 0
	for v := range resultchan {
		sum += v
	}
	return sum
}

// 通过done channel同步等待所有工作goroutine结束，然后关闭结果chan
func CloseResult(done chan struct{}, resultchan chan int, workers int) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(resultchan)
}

// 读取task chan并分发到worker goroutine处理，总数是 workers
func DistributeTask(taskchan <-chan task, workers int, done chan struct{}) {

	for i := 0; i < workers; i++ {
		go ProcessTask(taskchan, done)
	}

}

// 工作goroutine处理具体工作，并将处理结果发送结果 chan
func ProcessTask(taskchan <-chan task, done chan struct{}) {
	for t := range taskchan {
		t.do()
	}
	done <- struct{}{}
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
