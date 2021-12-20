package main

import (
	"filelistingserver/handler"
	"log"
	"net/http"
	"os"
)

/**
 * 定义了一个方法类型
 */
type appHandler func(
	writer http.ResponseWriter,
	request *http.Request) error

/**
 * 把会返回error的一个func，包装成一个不会返回error的方法
 * error在这个包装方法中处理
 */
func errWrapper(handler appHandler) func(
	writer http.ResponseWriter, request *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {

		// recover 处理panic
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v \n", r)
				http.Error(
					writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		// 处理 error
		err := handler(writer, request)
		if err != nil {

			log.Println("Error handling request. ", err.Error())

			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {

	// http://localhost:8888/list/web.go

	http.HandleFunc("/", errWrapper(handler.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
