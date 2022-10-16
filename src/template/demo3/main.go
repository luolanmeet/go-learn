package main

import (
	"html/template"
	"net/http"
	"os"
)

func tmpl(w http.ResponseWriter, r *http.Request) {
	path, _ := os.Getwd()
	path = path + "\\src\\template\\demo3\\"
	// 添加的顺序会有影响，使用了 template 的页面要先添加
	t, err := template.ParseFiles(path+"index.html", path+"header.html", path+"footer.html")
	if err != nil {
		panic(err)
	}

	err2 := t.Execute(w, nil)
	if err != nil {
		panic(err2)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/tmpl", tmpl)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
