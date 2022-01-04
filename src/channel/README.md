## channel
channel用于goroutine和goroutine之间交换数据
CSP模型
Don't communicate by sharing memory;share memory by communicating
不要通过共享内存来通信；通过通信来共享内存。

## 问题
fatal error: all goroutines are asleep - deadlock!

非缓冲的channel，在没有准备好的接受者(v := <-c )， 便开始发送数据(c <- 1)，编译时会提示死锁。
因为发送数据时(c <- 1)，协程是阻塞的，数据没被接受完成时不会被唤醒。
因此当没有接受者而开始发送数据时，必然出现死锁。

因此！有发送必须有接收

## 注意事项
channel，发数据和收数据，都会阻塞
channel的close，一定是发送方去close。

## 传统的同步机制
Waitgroup
Mutex
Cond

## select 调度
nil channel的使用

## 调度器
time.After、time.Tick