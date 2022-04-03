package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userinfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//password := c.Query("password")
		//u1 := userinfo{
		//	username: username,
		//	password: password,
		//}
		var u1 userinfo //结构体初始化
		err := c.ShouldBind(&u1)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
		fmt.Printf("%#v\n", u1)
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "ok",
		//})
	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/form", func(c *gin.Context) {
		var u1 userinfo //结构体初始化
		err := c.ShouldBind(&u1)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
		fmt.Printf("%#v\n", u1)
	})

	r.POST("/json", func(c *gin.Context) {
		var u1 userinfo //结构体初始化
		err := c.ShouldBind(&u1)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
		fmt.Printf("%#v\n", u1)
	})

	r.Run(":9000")
}
