package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	e := gin.Default()
	path, _ := os.Getwd()
	e.LoadHTMLGlob(path + "/demo2/templates/*")

	// locahost:8080/login
	e.GET("/login", Login)
	// locahost:8080/register
	e.GET("/register", Register)
	e.POST("login", DoLogin)
	e.POST("/register", DoRegister)
	e.Run()
}

func DoRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 获取数组参数
	hobby := c.PostFormArray("hobby")
	gender := c.PostForm("gender")
	city := c.PostForm("city")
	c.String(200, "username:%s password:%s hobby:%s gender:%s city:%s", username, password, hobby, gender, city)
}

func Register(c *gin.Context) {
	// 跳转到注册页面
	c.HTML(200, "register.html", nil)
}

func DoLogin(c *gin.Context) {
	// 从表单中获取数据
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.HTML(200, "welcome.html", gin.H{
		"username": username,
		"password": password,
	})
}

func Login(c *gin.Context) {
	// 跳转到登录页面
	c.HTML(200, "login.html", nil)
}
