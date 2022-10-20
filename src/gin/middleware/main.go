package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	// 默认使用了两个中间件 engine.Use(Logger(), Recovery())
	engine := gin.Default()

	// 添加中间件
	engine.Use(middlewareOne, middlewareTwo)

	// http://localhost:8080/testmw
	engine.GET("testmw", TestWM)
	engine.Run()
}

func middlewareOne(c *gin.Context) {
	fmt.Println("middlewareOne")
}

func middlewareTwo(c *gin.Context) {
	fmt.Println("middlewareTwo")
}

func TestWM(c *gin.Context) {
	c.String(200, "hello %s", "world")
}
