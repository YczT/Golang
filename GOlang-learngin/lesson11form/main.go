package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	// /login post请求  一次请求对应一次响应
	r.POST("/login", func(c *gin.Context) {
		//获取form表单提交的数据 ---1
		//username := c.PostForm("username")
		//password := c.PostForm("password")
		//获取form表单提交的数据 ---2
		//username := c.DefaultPostForm("username", "somebody")
		//password := c.DefaultPostForm("password", "some***")
		username, ok := c.GetPostForm("username")
		if !ok {
			username = "sb"
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "**"
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})

	r.Run(":9000")
}
