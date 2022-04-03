package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) { //获取个人信息之类的数据
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	r.POST("/index", func(c *gin.Context) { //注册账户
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/index", func(c *gin.Context) { //更新数据,修改数据
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/index", func(c *gin.Context) { //删除数据
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	//定义404页面
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "www.chenyan.com",
		})
	})
	//视频的首页和详情页
	//r.GET("/video/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
	//})
	//r.GET("/video/xx", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
	//})
	//r.GET("/video/oo", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
	//})
	//路由组的组  多用于区分不同的业务线或API版本
	//视频的首页和详情页（业务线）
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
		})
		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/xx"})
		})
		videoGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/oo"})
		})
	}
	//商城的首页和详情页（业务线）
	r.GET("/shop/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "/shop/index"})
	})
	r.GET("/shop/xx", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "/shop/xx"})
	})
	r.GET("/shop/oo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "/shop/oo"})
	})
	r.Run(":9000")
}
