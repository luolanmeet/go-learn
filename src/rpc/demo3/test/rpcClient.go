package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("Dialing:", err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn)) // 使用json格式解码

	var reply string
	err = client.Call("HelloService.Hello", "world", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)

	// 执行 nc -l 1234  可以看到客户端传输的数据为json格式

}
