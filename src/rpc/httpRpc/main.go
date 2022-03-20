package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc/httpRpc/hello"
)

func main() {

	rpc.RegisterName("HelloService", new(hello.HelloService))

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {

		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			Writer:     w,
			ReadCloser: r.Body,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":1234", nil)

	/**
	http的方式提供rpc服务

	 curl localhost:1234/jsonrpc -X POST \
		--data '{"method":"HelloService.Hello","params":["hello"],"id":0}'
	*/

}
