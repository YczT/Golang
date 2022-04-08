package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//GORM框架
//Object Relational Mapping
//对象(go中的结构体实例) 关系(关系型数据库) 映射
//帮忙写sql语句
//优点：提高开发效率
//缺点：牺牲执行性能，牺牲灵活性，弱化SQL能力

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	dsn := "root:CHENyan123@tcp(127.0.0.1:3306)/dbgorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//创建表，自动迁移(把结构体和数据表进行对应)
	db.AutoMigrate(&UserInfo{})
	u1 := UserInfo{1, "阿言", "男", "篮球"}
	u2 := UserInfo{2, "阿彤", "女", "看剧"}

	db.Create(&u1)
	db.Create(&u2)

	//var u UserInfo
	//var us []UserInfo
	//db.First(&u) //第一种查询方式，查询有效，因为目标struct是指针

	//result := map[string]interface{}{} //第二种查询方式，配合 Take 有效
	//db.Table("us").Take(&result)

	//result := map[string]interface{}{}// 无效
	//db.Table("us").First(&result)

	//result := map[string]interface{}{} //第三种查询方式，有效，因为通过 `db.Model()` 指定了 model
	//db.Model(&UserInfo{}).First(&result)

	//db.Model(&u).Update("hobby", "唱歌") //更新数据
	//
	//db.Model(&UserInfo{}).First(&result)
	//删除数据
	//db.Delete(&UserInfo{}, 2)

	//allrecord := db.Find(&UserInfo{})
	//record := allrecord.RowsAffected
	//fmt.Printf("\n%v", record)
}
