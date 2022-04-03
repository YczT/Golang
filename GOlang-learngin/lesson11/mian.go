package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//GET请求 URL ？后面的是querystring参数
	//key=value模式，多个key-value用 & 连接
	//例： web?query=阿言&age=24
	r.GET("/web", func(c *gin.Context) {
		//获取浏览器携带的querystring参数
		name := c.Query("query") //通过query获取请求中携带的query参数
		age := c.Query("age")
		//name := c.DefaultQuery("query", "some") //	取不到就用指定的默认值
		//name, nok := c.GetQuery("query") //取到值返回（值，true），取不到值返回（""，false）
		//if !nok {
		//	name = "somebody"
		//}
		//age, aok := c.GetQuery("age")
		//if !aok {
		//	age = "20"
		//}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":9000")
}
