package main

import (
	"net"
	"net/rpc"
	"rpc/reverseRpc/hello"
	"time"
)

func main() {

	rpc.RegisterName("HelloService", new(hello.HelloService))

	// 反向的rpc，反向代理，内网无法从外网链接，
	// 因此从内网发起连接，
	// 外网与内网建立连接后，外网再调用内网的rpc服务

	// 循环地去连接外网服务
	for {
		conn, _ := net.Dial("tcp", "127.0.0.1:1234")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}
		rpc.ServeConn(conn)
		conn.Close()
	}

}
