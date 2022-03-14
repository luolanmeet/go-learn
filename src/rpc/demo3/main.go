package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc/demo1/hello"
)

func main() {
	rpc.RegisterName("HelloService", new(hello.HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) // 使用 json格式编码

	// echo -e '{"method":"HelloService.Hello","params":["hello"],"id":1}' | nc localhost 1234
	// 可以看到服务端返回的是json格式的数据
}
