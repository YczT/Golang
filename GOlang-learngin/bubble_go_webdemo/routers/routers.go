package routers

/*
路由层
*/

import (
	"bubble_go_webdemo/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 连接前端
	r := gin.Default()
	//告诉gin框架模板文件引用的静态文件去哪里找
	//参数1：需求路径，参数2：我的项目中的哪个文件
	r.Static("/static", "static")
	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHanlder)

	//v1
	v1Group := r.Group("v1") //路由组
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateTodoById)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteTodoById)
	}
	return r
}
