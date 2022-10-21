package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()
	// 博客
	blog := r.Group("/blog")
	// localhost/blog/list
	{
		blog.POST("/list", F1)
		blog.POST("/add", F2)
		blog.POST("/delete", F3)
	}

	// 视频
	video := r.Group("/video")
	// localhost/video/list
	{
		video.POST("/list", F4)
		video.POST("/add", F5)
		video.POST("/delete", F6)
	}

}

func F1(c *gin.Context) {}
func F2(c *gin.Context) {}
func F3(c *gin.Context) {}
func F4(c *gin.Context) {}
func F5(c *gin.Context) {}
func F6(c *gin.Context) {}
