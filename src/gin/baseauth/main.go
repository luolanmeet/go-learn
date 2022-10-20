package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户信息，应该要是数据库中的数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@bar.com", "phone": "333"},
	"lena":   gin.H{"email": "lena@bar.com", "phone": "66633"},
}

func main() {

	engine := gin.Default()
	// 分组请求
	authorized := engine.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "4444",
	}))

	// http://localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {

		// 获取当前用户
		user := c.MustGet(gin.AuthUserKey).(string)
		fmt.Println(user)

		// 判断用户是否存在
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	engine.Run()
}
