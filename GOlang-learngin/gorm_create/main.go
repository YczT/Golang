package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int64
}

func (User) TableName() string {
	return "UserSelect"
}
func main() {
	dsn := "root:CHENyan123@tcp(127.0.0.1:3306)/dbgorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})

	//u1 := User{Name: "阿言", Age: 24}
	//db.Create(&u1)
	//u2 := User{Name: "阿彤", Age: 23}
	//db.Create(&u2)
	//u3 := User{Name: "啊狗", Age: 55}
	//u4 := User{Name: "啊猪", Age: 62}
	//db.Create(&u3)
	//db.Create(&u4)
	//查询
	var user User   //声明模型结构体类型变量
	db.First(&user) //传指针（传入的变量地址，然后通过地址修改变量）只有主键是数字才适用
	//fmt.Printf("第一个用户:%v\n", user)

	var users []User
	db.Debug().Find(&users)
	//fmt.Printf("全部用户:%v\n", users)
	//where条件查询(sql语句拼接)
	var usertong User
	var usertong2 []User
	var usertong3 User

	db.Where("name=?", "阿彤").Find(&usertong)
	fmt.Printf("名字叫阿彤的用户:%v\n", usertong)
	db.Where("name <> ?", "阿言").Find(&usertong2)
	fmt.Printf("名字叫阿彤的用户:%v\n", usertong2)
	db.Where("name=? AND age>=?", "阿言", 22).Find(&usertong3)
	fmt.Printf("名字叫阿彤的用户:%v\n", usertong3)

	//通过结构体或map作为where条件查询
	var user_struct User
	var user_map User
	db.Where(&User{Name: "阿言", Age: 24}).Find(&user_struct)
	db.Where(map[string]interface{}{"Name": "阿彤", "Age": 23}).Find(&user_map)
	//fmt.Printf("结构体或map作为条件查询：%v|%v\n", user_struct, user_map)
	//	当使用结构作为条件查询时，GORM 只会查询非零值字段。这意味着如果您的字段值为 0、''、false 或其他零值，该字段不会被用于构建查询条件
	//	如果想要包含零值查询条件，你可以使用 map，其会包含所有 key-value 的查询条件
	//  Limit 指定获取记录的最大数量 Offset 指定在开始返回记录之前要跳过的记录数量
	db.Limit(2).Offset(1).Find(&users)
	fmt.Printf("%v\n", users)
	
}
