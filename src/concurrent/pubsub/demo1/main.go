package main

import (
	"concurrent/pubsub/demo1/pubsub"
	"fmt"
	"strings"
	"time"
)

func main() {

	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	// 订阅全部主题
	all := p.Subscribe()
	// 订阅golang
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})
	p.Publish("hello world")
	p.Publish("hello golang")
	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()
	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()

	time.Sleep(3 * time.Second)
}
