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
	return "userupdate"
}

func main() {
	dsn := "root:CHENyan123@tcp(127.0.0.1:3306)/dbgorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	//u1 := User{Name: "阿言", Age: 24, Active: true}
	//u2 := User{Name: "阿彤", Age: 23, Active: false}
	//u3 := User{Name: "阿猪", Age: 40, Active: true}
	//db.Create(&u1)
	//db.Create(&u2)
	//db.Create(&u3)

	var users []User
	db.Find(&users)
	fmt.Printf("全部记录:%v\n", users)

	var user User
	db.Find(&user)
	fmt.Printf("第一条记录:%v\n", user)
	//更新
	//user.Name = "小哥哥真帅"
	//user.Age = 28
	//db.Debug().Save(&user) //默认会修改所有字段(把所有字段从数据库拿到后端)
	//fmt.Printf("第一条修改记录:%v\n", user)

	//m1 := map[string]interface{}{
	//	"name":   "chenyanyan",
	//	"age":    29,
	//	"active": true,
	//}

	//db.Model(&user).Updates(m1) //map里列出来的所有字段都会更新
	//fmt.Printf("第一条修改记录第二次:%v\n", user)
	//db.Model(&user).Select("age").Updates(m1) //只更新age字段
	//fmt.Printf("第一条修改记录第三次:%v\n", user)
	//db.Model(&user).Omit("active").Updates(m1)
	//fmt.Printf("第一条修改记录第四次:%v\n", user)
	// 批量更新时，钩子函数不会触发(钩子函数：在CRUD之前或之后进行的操作)
	// 将表中所有的用户年龄 +2
	db.Model(&users).Updates(map[string]interface{}{"age": gorm.Expr("age + ?", 2)})
	fmt.Printf("将所有用户的年龄加2 :%v\n", &users)

}
