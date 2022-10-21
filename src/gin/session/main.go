package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	// gin 没有提供 session 的能力，需要再依赖其他包
	// 其他包的实现方式，就是通过提供 gin 的中间件的方式实现

	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))

	// 注入中间件！
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", func(c *gin.Context) {

		session := sessions.Default(c)
		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Save()
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})

	r.Run()
}
