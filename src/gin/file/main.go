package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {

	r := gin.Default()

	path, _ := os.Getwd()
	r.LoadHTMLGlob(path + "/file/templates/*")

	r.GET("/upload", Upload)
	r.POST("/upload", DoUpload)

	r.Run()
}

func DoUpload(c *gin.Context) {

	// 单文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// 上传文件到项目根目录，使用原文件名
	c.SaveUploadedFile(file, file.Filename)
	c.String(http.StatusOK, fmt.Sprintf("%s uploaded", file.Filename))
}

func Upload(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
