package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"net/http"
	"strings"
)

func httpService(w http.ResponseWriter, r *http.Request) {

	// 解析参数，默认不会解析
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println(k, strings.Join(v, ""))
	}

	if result := r.Form["success"]; result != nil && result[0] == "false" {
		fmt.Fprintf(w, "<?xml version=\"1.0\" encoding=\"utf-8\"?><response><flag>false</flag><code>500</code><message></message></response>")
	} else {
		fmt.Fprintf(w, "<?xml version=\"1.0\" encoding=\"utf-8\"?><response><flag>success</flag><code>200</code><message></message></response>")
	}

}

func main() {
	// 设置访问的路由
	http.HandleFunc("/", httpService)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
