package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc/contextRpc/hello"
)

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			defer conn.Close()
			p := rpc.NewServer()
			p.Register(&hello.HelloService{Conn: conn})
			p.ServeConn(conn)
		}()
	}

}
