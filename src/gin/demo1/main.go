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

func TestQueryString(c *gin.Context) {
	// 获取 url 中的请求数据
	username := c.Query("username")
	age := c.DefaultQuery("age", "50")
	c.String(200, "username:%s, age:%s", username, age)
}

func TestPostParam(c *gin.Context) {
	// 从表单中获取数据
	username := c.PostForm("username")
	age := c.DefaultPostForm("age", "123")
	c.String(200, "username:%s, age:%s", username, age)
}

func TestPathParam(c *gin.Context) {
	// restful 风格的 api 获取请求参数
	username := c.Param("username")
	age := c.Param("age")
	c.String(200, "username:%s, age:%s", username, age)
}

func main() {
	engine := gin.Default()
	engine.GET("/hello", Hello)
	engine.GET("/ping", Ping)
	// localhost:8080/testQueryString?username=cck&age=50
	engine.GET("/testQueryString", TestQueryString)
	// localhost:8080/testPathParam/cck/50
	engine.GET("/testPathParam/:username/:age", TestPathParam)
	engine.Run()
}
