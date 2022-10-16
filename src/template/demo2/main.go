package main

import (
	"net/http"
	"os"
	"text/template"
)

func tmpl(w http.ResponseWriter, r *http.Request) {
	path, _ := os.Getwd()
	t, err := template.ParseFiles(path + "\\src\\template\\demo2\\index.html")
	if err != nil {
		panic(err)
	}

	data := make(map[string]interface{})
	data["Hello"] = "hello world"
	data["Friend"] = []string{"as", "ac", "hg", "ft", ".."}

	t.Execute(w, data)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/tmpl", tmpl)
	server.ListenAndServe()
}
