package main

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	// 返回字符串
	c.String(200, "hello %s", "world")
}

func Ping(c *gin.Context) {
	// 返回 JSON 格式数据
	// gin.H 是 gin 里边 map 的一个快捷方式
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	engine := gin.Default()
	engine.GET("/hello", Hello)
	engine.GET("/ping", Ping)
	engine.Run()
}
