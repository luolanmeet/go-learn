package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc/watch/store"
)

func main() {

	rpc.RegisterName("KVStoreService", store.NewKVStoreService)
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
