package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
}

func (User) TableName() string {
	return "userdelete"
}

func main() {

	dsn := "root:CHENyan123@tcp(127.0.0.1:3306)/dbgorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	//u1 := User{Name: "ayan", Age: 24, Active: true}
	//u2 := User{Name: "atong", Age: 23, Active: false}
	//db.Create(&u1)
	//db.Create(&u2)
	// 删除

	var u = User{}
	u.ID = 1
	//u.Name = "ayan"
	//	在删除记录的时候需要指定主键，(用u.ID做删除就只删除u.ID所指定的记录) 否则会删除全部记录
	//	(用u.Name(非主键)做删除就删除了全部的记录) 要么就带上where条件
	//db.Debug().Delete(&u)

	//如果您的模型包含了一个 gorm.deletedat 字段（gorm.Model 已经包含了该字段)，它将自动获得软删除的能力！
	//拥有软删除能力的模型调用 Delete 时，记录不会被数据库。但 GORM 会将 DeletedAt 置为当前时间， 并且你不能再通过普通的查询方法找到该记录。
	var users []User
	db.Find(&users)
	fmt.Printf("全部记录:%v\n", &users)

	//	 可以通过Unscoped查看被软删除的字段，也可以通过Unscoped永久删除字段
	var user User
	db.Unscoped().Where("age = 24").Find(&user)
	fmt.Printf("被软删除的字段：%v\n", &user)

	db.Unscoped().Delete(&u)
	db.Find(&users)
	fmt.Printf("全部记录:%v\n", &users)
}
