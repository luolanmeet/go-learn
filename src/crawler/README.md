go get golang.org/x/text

go get golang.org/x/net/html


> tag 1.0.0 单任务版;
> 
> tag 2.0.0 并发版;
> 每个request一个goroutine;
> 所有worker争抢一个channel;
> 
> tag 3.0.0 并发版;
> schedule维护一个request和worker的队列;
> 