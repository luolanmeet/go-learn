package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {

	engine := gin.Default()
	path, _ := os.Getwd()

	// 引用静态文件
	engine.Static("/assets", path+"/demo3/assets")
	engine.LoadHTMLGlob(path + "/demo3/templates/*")

	engine.GET("/gostatic", GoStatic)
	engine.Run()
}

func GoStatic(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
