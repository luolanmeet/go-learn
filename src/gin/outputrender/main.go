package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()

	// 支持将输出为不同的格式
	r.GET("/test", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"name": "cck",
			"age":  50,
		})

		c.XML(200, gin.H{
			"name": "cck",
			"age":  50,
		})

		c.HTML(200, "login.html", nil)

		c.String(200, "hello world")
		//c.ProtoBuf(200, nil)
	})

}
