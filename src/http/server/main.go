package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func testHttpServer() {
	f := func(resp http.ResponseWriter, req *http.Request) {
		io.WriteString(resp, "hello world")
	}

	// 相应路径，需要带 /
	http.HandleFunc("/hello", f)
	// 设置监听端口，需要带 :
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type countHandler struct {
	mu sync.Mutex
	n  int
}

// 相当于 countHandler 实现了 handler 接口的方法
func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func testHttpServer2() {
	http.Handle("/count", new(countHandler))
	http.ListenAndServe(":8888", nil)
}

func main() {
	//testHttpServer2()
	testHttpServer2()
}
