package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc/standardInterfaceRpc/hello"
)

func main() {
	err := hello.RegisterHelloService(new(hello.HelloService))
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}

}
