package main

import (
	"bubble_go_webdemo/dao"
	"bubble_go_webdemo/models"
	"bubble_go_webdemo/routers"
)

func main() {
	//连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	r := routers.SetupRouter()
	r.Run(":8999")
}
