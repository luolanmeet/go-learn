package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	e := gin.Default()
	path, _ := os.Getwd()
	e.LoadHTMLGlob(path + "/demo2/templates/*")

	// localhost:8080/login
	e.GET("/login", Login)
	// localhost:8080/register
	e.GET("/register", Register)
	// localhost:8080/register2
	e.GET("/register2", Register2)

	e.POST("login", DoLogin)
	e.POST("/register", DoRegister)
	e.POST("/register2", DoRegister2)
	err := e.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Login(c *gin.Context) {
	// 跳转到登录页面
	c.HTML(200, "login.html", nil)
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

func Register(c *gin.Context) {
	// 跳转到注册页面
	c.HTML(200, "register.html", nil)
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

type User struct {
	// form 表单标签
	Username string   `form:"username"`
	Password string   `form:"password"`
	Hobby    []string `form:"hobby"`
	Gender   string   `form:"gender"`
	City     string   `form:"city"`
}

func Register2(c *gin.Context) {
	// 跳转到注册页面
	c.HTML(200, "register2.html", nil)
}

func DoRegister2(c *gin.Context) {
	// 用结构体接收参数
	var user User
	// 可绑定表单的数据、url中的数据 xxx?username=123，结构体字段的标签是 form
	// 如要绑定 path 中的参数 /login/:username，需要换 ShouldBindUri api，结构体字段的标签是 uri
	c.ShouldBind(&user)
	c.String(200, "User:%s", user)
}
