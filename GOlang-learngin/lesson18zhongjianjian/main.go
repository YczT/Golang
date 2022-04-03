package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHeandler(c *gin.Context) {
	fmt.Println("index in ...")
	name, ok := c.Get("name") //在上下文取值（跨中间件存取值）
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}
func shopHeandler(c *gin.Context) {
	fmt.Println("shop in ...")
	c.JSON(http.StatusOK, gin.H{
		"msg": "shop",
	})
}
func vedioHeandler(c *gin.Context) {
	fmt.Println("vedio in ...")
	c.JSON(http.StatusOK, gin.H{
		"msg": "vedio",
	})
}

//定义一个中间件m1：统计请求处理函数的耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	//计时
	start := time.Now()
	//go funcxx(c.Copy()) //启动一个新的goroutine的函数funcxx中只能使用上下文‘c’的拷贝，必须使用只读对象
	c.Next() //调用后续的处理函数 indexHeandler
	//c.Abort() //阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out...")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	c.Set("name", "constin") //在上下文中设置值
	//c.Next() //调用后续的处理函数 indexHeandler
	//c.Abort() //阻止调用后续的处理函数
	//return
	fmt.Println("m2 out...")
}

func autoMiddleware(docheck bool) gin.HandlerFunc {
	//连接数据库或一些其他准备工作
	return func(c *gin.Context) {
		if docheck {
			//存放具体逻辑
			//是否登录的判断
			//	if是登录用户
			//	c.Next()
			//	else
			//	c.Abort()
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()                   //默认使用了Logger()和Recover()中间件
	r.Use(m1, m2, autoMiddleware(false)) //全局注册中间件
	r.GET("/index", indexHeandler)
	r.GET("/shop", shopHeandler)
	r.GET("/vedio", vedioHeandler)

	r.Run(":8999")
}
