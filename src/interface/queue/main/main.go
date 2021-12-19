package main

import (
	"fmt"
	"interface/queue"
)

func main() {
	queue := queue.Queue{}
	queue.Push("abc")
	fmt.Println(queue.Pop())
	fmt.Println(queue.IsEmpty())
}
