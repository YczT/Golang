package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"status": "ok",
		//})
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com") //路由重定向
	})
	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b" //把请求的URI修改
		r.HandleContext(c)        //继续后续的处理
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})

	r.Run(":9000")
}
