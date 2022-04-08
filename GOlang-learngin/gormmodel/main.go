package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//type User struct {
//	gorm.Model
//	Name         string
//	Age          sql.NullInt64 `gorm:"column:user_age"` //零值类型
//	Birthday     *time.Time
//	Email        string  `gorm:"type:varchar(120);unique_index"` //类型，唯一索引（不重复）
//	Rola         string  `gorm:"size:255"`                       //字段大小
//	MemberNumber *string `gorm:"unique;not null"`                //唯一且不为空
//	Num          int     `gorm:"AUTO_INCREMENT"`                 //自增
//	Address      string  `gorm:"index:addr"`                     //创建名为addr的索引
//	IgnoreMe     int     `gorm:"-"`                              //忽略本字段
//}

//1 定义类型

type User struct {
	ID   int64
	Name sql.NullString `gorm:"default:美丽的大叔"`
	Age  int64
}

func (User) TableName() string {
	return "createTab"
}

func main() {
	dsn := "root:CHENyan123@tcp(127.0.0.1:3306)/dbgorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, //创建表的时候不更改表名（user，不是users）
	})
	if err != nil {
		panic(err)
	}
	//2 把模型与数据库中的表对应起来
	db.AutoMigrate(&User{}) //自动迁移（创建表）

	//3 创建数据
	u := User{Name: sql.NullString{String: "", Valid: true}, Age: 27}
	db.Debug().Create(&u)

}
