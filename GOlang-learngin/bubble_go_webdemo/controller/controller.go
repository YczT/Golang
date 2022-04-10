package controller

import (
	"bubble_go_webdemo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
url --> controller --> logic --> model
请求来了  -->  控制器  -->  业务逻辑  -->  模型层的增删改查
*/

func IndexHanlder(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	//前端页面点击待办事项，点击提交发请求到这里
	// 1.从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	// 2.存入数据库
	// 3.返回响应
	err := models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(c *gin.Context) {
	// 查询表中的所有数据
	todolist, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}

func UpdateTodoById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "Id不存在"})
	}
	todoupdate, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	c.BindJSON(&todoupdate)
	if err = models.UpdateATodo(todoupdate); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoupdate)
	}
}
func DeleteTodoById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "Id不存在"})
	}
	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
	}
}
