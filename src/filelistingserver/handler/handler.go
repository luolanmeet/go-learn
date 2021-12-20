package handler

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

/**
 * 用一个字符串类型，来实现接口 Error 和 Message接口
 */
type userError string

func (err userError) Error() string {
	return err.Message()
}

func (err userError) Message() string {
	return string(err)
}

func HandleFileList(
	writer http.ResponseWriter,
	request *http.Request) error {

	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must start with " + prefix)
	}

	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}
