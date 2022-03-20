package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("Dialing:", err)
	}

	/**
	同步调用的方式
	*/
	//var reply string
	//err = client.Call("HelloService.Hello", "world", &reply)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(reply)

	/**
	异步调用的方式
	*/
	helloCall := client.Go("HelloService.Hello", "world", new(string), nil)
	helloCall = <-helloCall.Done
	if err := helloCall.Error; err != nil {
		log.Fatal(err)
	}
	args := helloCall.Args.(string)
	reply := helloCall.Reply.(*string)
	fmt.Println(args, *reply)
}
