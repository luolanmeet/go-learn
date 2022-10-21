package main

import "github.com/gin-gonic/gin"

func Handler(c *gin.Context) {
	// 获得 cookie
	cookie, err := c.Cookie("username")
	if err != nil {

		cookie = "cck"
		// 设置 cookie
		// cookie 名称、值、最大存活时间、路径、域名、secure、httpOnly
		// Secure 属性为 true 时，cookie 只能用 https 协议发送给服务器，用 http 协议无法发送。
		// HttpOnly 属性 为 true 时，cookie 不能被 js 获取到，无法用 document.cookie 获取到 cookie 的内容。
		c.SetCookie("username", cookie, 60*60, "/", "localhost", false, true)
	}
	c.String(200, "测试 cookie")
}

func main() {
	engine := gin.Default()
	engine.GET("/cookie", Handler)
	engine.Run()
}
