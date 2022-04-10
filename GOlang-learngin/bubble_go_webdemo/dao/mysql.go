package dao

/*
dao:Data Access Object(数据访问对象)
*/

import (
	mysql2 "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMysql() (err error) {
	dsn := "root:CHENyan123@tcp(127.0.0.1:3306)/dbgoweb?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql2.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	return
}
